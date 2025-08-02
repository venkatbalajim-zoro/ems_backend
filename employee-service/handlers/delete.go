package handlers

import (
	"employee-service/configs"
	"employee-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Delete(context *gin.Context) {
	var input models.Input

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to fetch the input data",
		})
		return
	}

	result := configs.Database.Table("employees").Where("employee_id = ?", input.EmployeeID).Delete(&models.Employee{})
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error occured",
		})
		return
	} else if result.RowsAffected == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"error": "Employee details not found",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Employee details are deleted successfully ...",
	})
}
