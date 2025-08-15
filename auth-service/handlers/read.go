package handlers

import (
	"auth-service/configs"
	"auth-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Read(context *gin.Context) {
	var accounts []models.Account

	err := configs.Database.Table("accounts").Find(&accounts).Error

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to fetch the accounts data from the database.",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Accounts data is fetched successfully.",
		"data":    accounts,
	})
}
