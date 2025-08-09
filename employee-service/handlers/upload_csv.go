package handlers

import (
	"employee-service/configs"
	"employee-service/models"
	"encoding/csv"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func UploadCSV(context *gin.Context) {
	var input map[string]string
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to fetch the file path",
		})
		return
	}
	path := input["path"]

	file, err := os.Open(path)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to read the file",
		})
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to read the content",
		})
		return
	}

	if len(records) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Empty CSV file",
		})
		return
	}

	var data []models.Employee
	for _, record := range records {
		id, err := strconv.Atoi(record[0])
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid data in the CSV file",
			})
			return
		}
		firstName, lastName, email, phone, gender := record[1], record[2], record[3], record[4], record[5]
		deptId, err := strconv.Atoi(record[6])
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid data in the CSV file",
			})
			return
		}
		designation := record[7]
		salary, err := strconv.ParseFloat(record[8], 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid data in the CSV file",
			})
			return
		}

		hireDate, err := time.Parse(time.DateOnly, record[9])
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid data in the CSV file",
			})
			return
		}

		elt := models.Employee{
			EmployeeID:   id,
			FirstName:    firstName,
			LastName:     lastName,
			Email:        email,
			Phone:        phone,
			Gender:       gender,
			DepartmentID: deptId,
			Designation:  designation,
			Salary:       salary,
			HireDate:     hireDate,
		}

		data = append(data, elt)
	}

	err = configs.Database.Table("employees").Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "employee_id"}},
		DoUpdates: clause.AssignmentColumns([]string{
			"first_name", "last_name", "email", "phone",
			"gender", "department_id", "designation", "salary", "hire_date",
		}),
	}).Create(&data).Error

	if err != nil {
		context.JSON(http.StatusConflict, gin.H{
			"error": "Unable to upload CSV data in the database",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Data is uploaded successfully",
	})
}
