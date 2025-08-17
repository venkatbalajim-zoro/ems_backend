package handlers

import (
	"department-service/configs"
	"department-service/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

func isValidData(data models.Department) bool {
	return data.ID != 0 && data.Name != ""
}

func Add(context *gin.Context) {
	var data models.Department
	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to fetch the input data.",
		})
		return
	}

	if !isValidData(data) {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Some required data are missing.",
		})
		return
	}

	err := configs.Database.Table("departments").Create(&data).Error
	if err != nil {
		var sqlError *mysql.MySQLError
		if errors.As(err, &sqlError) {
			switch sqlError.Number {
			case 1062:
				context.JSON(http.StatusConflict, gin.H{
					"error": "Already a department exists with same ID.",
				})
				return
			default:
				context.JSON(http.StatusInternalServerError, gin.H{
					"error": "Internal server error occured.",
				})
				return
			}
		}

		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unexpected error while inserting department.",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Department details are added successfully.",
		"data":    data,
	})
}
