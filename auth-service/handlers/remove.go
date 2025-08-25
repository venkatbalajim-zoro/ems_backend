package handlers

import (
	"auth-service/configs"
	"auth-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Remove(context *gin.Context) {
	var input models.Input
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to fetch the input data.",
		})
		return
	}

	username := input.Username
	if username == "" {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Username must not be empty.",
		})
		return
	}

	result := configs.Database.Table("accounts").Where("username = ?", username).Delete(&models.Account{})

	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to delete the data in the database.",
		})
		return
	}

	if result.RowsAffected == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "Username not found to delete the data.",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Account is deleted successfully.",
	})
}
