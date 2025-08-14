package handlers

import (
	"employee-service/configs"
	"employee-service/models"
	"encoding/csv"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

func UploadCSV(context *gin.Context) {
	// Get uploaded file
	file, err := context.FormFile("file")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "CSV file is required",
		})
		return
	}

	// Open uploaded file
	src, err := file.Open()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to open uploaded file",
		})
		return
	}
	defer src.Close()

	// Read CSV content
	reader := csv.NewReader(src)
	records, err := reader.ReadAll()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to read CSV content",
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
		if len(record) != 10 {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "Each row must have exactly 10 fields",
			})
			return
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Employee ID"})
			return
		}

		deptId, err := strconv.Atoi(record[6])
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Department ID"})
			return
		}

		salary, err := strconv.ParseFloat(record[8], 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Salary"})
			return
		}

		data = append(data, models.Employee{
			EmployeeID:   id,
			FirstName:    record[1],
			LastName:     record[2],
			Email:        record[3],
			Phone:        record[4],
			Gender:       record[5],
			DepartmentID: deptId,
			Designation:  record[7],
			Salary:       salary,
			HireDate:     record[9],
		})
	}

	// Insert or update (UPSERT)
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
		"message": "Data uploaded successfully",
	})
}
