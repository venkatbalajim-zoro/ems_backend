package handlers

import (
	"employee-service/configs"
	"employee-service/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

func Add(context *gin.Context) {
	var data models.Employee
	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to fetch the input data",
		})
		return
	}

	err := configs.Database.Create(&data).Error
	if err != nil {
		var sqlError *mysql.MySQLError
		if errors.As(err, &sqlError) {
			switch sqlError.Number {
			case 1452:
				context.JSON(http.StatusBadRequest, gin.H{
					"error": "The department ID does not exist.",
				})
				return
			case 1062:
				context.JSON(http.StatusBadRequest, gin.H{
					"error": "Already an employee exists with same ID",
				})
				return
			case 1048:
				context.JSON(http.StatusConflict, gin.H{
					"error": "Data must not be null",
				})
				return
			case 3819:
				context.JSON(http.StatusConflict, gin.H{
					"error": "Data is violating the check constraints in the database",
				})
				return
			default:
				context.JSON(http.StatusInternalServerError, gin.H{
					"error": "Internal server error while inserting employee details.",
				})
				return
			}
		}

		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unexpected error while inserting employee.",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Employee details are added successfully.",
		"data":    data,
	})
}
