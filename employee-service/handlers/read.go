package handlers

import (
	"employee-service/configs"
	"employee-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Read(context *gin.Context) {
	var rows []models.Employee

	err := configs.Database.Table("employees").Find(&rows).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error occured",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "All employees details are fetched successfully ...",
		"data":    rows,
	})
}
