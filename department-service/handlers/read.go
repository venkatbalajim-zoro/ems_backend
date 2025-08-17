package handlers

import (
	"department-service/configs"
	"department-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Read(context *gin.Context) {
	var rows []models.Department

	err := configs.Database.Table("departments").Find(&rows).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error occured.",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "All department details are fetched successfully.",
		"data":    rows,
	})
}
