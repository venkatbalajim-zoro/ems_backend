package handlers

import (
	"auth-service/configs"
	"auth-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Recovery(context *gin.Context) {
	var input models.Input

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input data",
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

	context.JSON(http.StatusOK, gin.H{
		"message":  "Password fetched successfully",
		"password": account.Password,
	})
}

// this feature will be improved like sending the password through email.
