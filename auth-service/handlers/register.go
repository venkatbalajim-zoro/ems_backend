package handlers

import (
	"auth-service/configs"
	"auth-service/models"
	"auth-service/utils"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

func Register(context *gin.Context) {
	var input models.Account
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to fetch the input data.",
		})
		return
	}

	if input.Username == "" || input.Password == "" || input.EmployeeID == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Please provide all the details.",
		})
		return
	}

	err := configs.Database.Table("accounts").Create(&input).Error
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				context.JSON(http.StatusConflict, gin.H{
					"error": "Username already exists.",
				})
				return
			case 1452:
				context.JSON(http.StatusConflict, gin.H{
					"error": "There is no employee using this ID. Please check the details.",
				})
				return
			default:
				context.JSON(http.StatusInternalServerError, gin.H{
					"error": "Unable to register the account details.",
				})
			}
			return
		}
	}

	var data struct{ Email string }
	err = configs.Database.Table("employees").Select("email").Where("employee_id = ?", input.EmployeeID).First(&data).Error
	if err != nil {
		log.Printf("Error in fetching the email from database: %s\n", err)
	} else {
		err = utils.SendEmail(models.Email{
			ToEmails:   []string{data.Email},
			Username:   input.Username,
			Password:   input.Password,
			EmployeeID: input.EmployeeID,
			Action:     "created",
			DateTime:   time.Now(),
		})

		if err != nil {
			log.Printf("Error in sending the email: %s\n", err)
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Account registered successfully.",
	})
}
