package handlers

import (
	"auth-service/configs"
	"auth-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Delete(context *gin.Context) {
	username, ok := context.Get("username")

	if !ok {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to fetch the username from the request.",
		})
		return
	}

	result := configs.Database.Table("accounts").Where("username = ?", username).Delete(&models.Account{})
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to delete your account from the database.",
		})
		return
	} else if result.RowsAffected == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "No account is registered with this username.",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Your account is deleted successfully.",
	})
}
