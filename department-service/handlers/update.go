package handlers

import (
	"department-service/configs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Update(context *gin.Context) {
	var data map[string]interface{}

	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to fetch the input data",
		})
		return
	}

	value, ok := data["id"]
	if !ok {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Department ID is required to update the details",
		})
		return
	}

	delete(data, "id")

	result := configs.Database.Table("departments").Where("id = ?", value).Updates(data)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unexpected error occured",
		})
		return
	} else if result.RowsAffected == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "No department found with this ID",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Department details are updated successfully ...",
	})
}
