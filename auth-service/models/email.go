package models

import "time"

type Email struct {
	ToEmails   []string  `json:"to_emails"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	EmployeeID int       `json:"employee_id"`
	Action     string    `json:"action"`
	DateTime   time.Time `json:"date_time"`
}
