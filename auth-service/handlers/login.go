package handlers

import (
	"auth-service/configs"
	"auth-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Login(context *gin.Context) {
	var input models.Input
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input data",
		})
		return
	}

	var account models.Account
	if err := configs.Database.Table("accounts").Where("username = ?", input.Username).First(&account).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(http.StatusUnauthorized, gin.H{
				"error": "No account is available with this username",
			})
			return
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
			return
		}
	}

	if account.Password != input.Password {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": "Please ensure your password",
		})
		return
	}

	token, err := configs.GenerateToken(account.Username, account.EmployeeID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to create the token for the authorization",
		})
		return
	} else {
		context.JSON(http.StatusOK, gin.H{
			"message": "Login successful",
			"token":   token,
		})
		return
	}
}
