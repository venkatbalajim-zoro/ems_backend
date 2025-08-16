package handlers

import (
	"auth-service/configs"
	"auth-service/models"
	"auth-service/utils"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Recovery(context *gin.Context) {
	var input models.Input

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to fetch the input data.",
		})
		return
	}

	var username = input.Username
	var account models.Account
	if err := configs.Database.Table("accounts").Where("username = ?", username).First(&account).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "This username is not registered in the database.",
		})
		return
	}

	var data struct{ Email string }
	err := configs.Database.Table("employees").Select("email").Where("employee_id = ?", account.EmployeeID).First(&data).Error
	if err != nil {
		log.Printf("Error in fetching the email from database: %s\n", err)
	} else {
		err = utils.SendEmail(models.Email{
			ToEmails:   []string{data.Email},
			Username:   account.Username,
			Password:   account.Password,
			EmployeeID: account.EmployeeID,
			Action:     "recovered",
			DateTime:   time.Now(),
		})

		if err != nil {
			log.Printf("Error in sending the email: %s\n", err)
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Password fetched successfully. Please check the email.",
	})
}
