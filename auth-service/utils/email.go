package utils

import (
	"auth-service/configs"
	"auth-service/models"
	"fmt"
	"net/smtp"
)

func SendEmail(details models.Email) error {
	host := configs.GetEnv("SMTP_HOST", "smtp.example.com")
	port := configs.GetEnv("SMTP_PORT", "587")
	from := configs.GetEnv("SMTP_EMAIL", "sample@example.com")
	password := configs.GetEnv("SMTP_PASSWORD", "password")

	address := fmt.Sprintf("%s:%s", host, port)
	auth := smtp.PlainAuth(from, from, password, host)
	to := details.ToEmails
	message := fmt.Sprintf(`Hello %s,
Your account has been successfully %s in the employee management system. 

Account details:
	- Username: %s
	- Password: %s
	- Employee ID: %d
	- Date-time: %s

Thank you,
HR department
	`, details.Username, details.Action, details.Username,
		details.Password, details.EmployeeID, details.DateTime.Format("02-01-2006 15:04:05"),
	)

	err := smtp.SendMail(address, auth, from, to, []byte(message))

	if err != nil {
		return err
	} else {
		return nil
	}
}
