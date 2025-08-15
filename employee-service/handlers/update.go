package handlers

import (
	"employee-service/configs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Update(context *gin.Context) {
	var input map[string]interface{}
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to fetch the input data.",
		})
		return
	}

	value, ok := input["employee_id"]
	if !ok {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Employee ID is required to update the details.",
		})
		return
	}

	delete(input, "employee_id")

	result := configs.Database.Table("employees").Where("employee_id = ?", value).Updates(input)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error occured.",
		})
		return
	} else if result.RowsAffected == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "No employee found with this ID.",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Employee details are updated successfully.",
	})
}
