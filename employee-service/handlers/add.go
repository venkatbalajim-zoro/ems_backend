package handlers

import (
	"employee-service/configs"
	"employee-service/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

func isValidData(e models.Employee) bool {
	if e.EmployeeID == 0 ||
		e.FirstName == "" ||
		e.LastName == "" ||
		e.Email == "" ||
		e.Phone == "" ||
		e.Gender == "" ||
		e.DepartmentID == 0 ||
		e.Designation == "" ||
		e.Salary == 0.0 ||
		e.HireDate.IsZero() {
		return false
	}
	return true
}

func Add(context *gin.Context) {
	var data models.Employee
	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to fetch the input data",
		})
		return
	}

	if !isValidData(data) {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing required fields",
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
