package handlers

import (
	"auth-service/configs"
	"auth-service/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

func Register(context *gin.Context) {
	var input models.Account
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input data",
		})
		return
	}

	if input.Username == "" || input.Password == "" || input.EmployeeID == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Please provide all the details",
		})
		return
	}

	err := configs.Database.Create(&input).Error
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) {
			switch mysqlErr.Number {
			case 1062:
				context.JSON(http.StatusConflict, gin.H{
					"error": "Username already exists",
				})
				return
			case 1452:
				context.JSON(http.StatusConflict, gin.H{
					"error": "There is no employee using this ID. Please check the details.",
				})
				return
			default:
				context.JSON(http.StatusInternalServerError, gin.H{
					"error": "Unable to register the account details",
				})
			}
			return
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Account registered successfully",
	})
}
