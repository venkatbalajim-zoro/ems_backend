package handlers

import (
	"department-service/configs"
	"department-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Delete(context *gin.Context) {
	var input map[string]int

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to fetch the input data",
		})
		return
	}

	result := configs.Database.Table("departments").Where("id = ?", input["id"]).Delete(&models.Department{})
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error occured",
		})
		return
	} else if result.RowsAffected == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "Department details not found",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Department details are deleted successfully ...",
	})
}
