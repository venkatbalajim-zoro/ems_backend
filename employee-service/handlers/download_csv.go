package handlers

import (
	"employee-service/configs"
	"employee-service/models"
	"encoding/csv"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func DownloadCSV(context *gin.Context) {
	// Fetch employees from database
	var records []models.Employee
	result := configs.Database.Table("employees").Find(&records)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to fetch employee data",
		})
		return
	}

	// Set headers for file download
	context.Header("Content-Description", "File Transfer")
	context.Header("Content-Disposition", `attachment; filename="employees.csv"`)
	context.Header("Content-Type", "text/csv")

	// Create CSV writer for HTTP response
	writer := csv.NewWriter(context.Writer)
	defer writer.Flush()

	// Write header row
	if err := writer.Write([]string{
		"Employee ID", "First Name", "Last Name", "Email ID", "Phone",
		"Gender", "Department ID", "Designation", "Salary", "Hire Date",
	}); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to write CSV header",
		})
		return
	}

	// Write data rows
	for _, record := range records {
		if err := writer.Write([]string{
			strconv.Itoa(record.EmployeeID),
			record.FirstName,
			record.LastName,
			record.Email,
			record.Phone,
			record.Gender,
			strconv.Itoa(record.DepartmentID),
			record.Designation,
			strconv.FormatFloat(record.Salary, 'f', 2, 64), // normal decimal format
			record.HireDate.Format(time.DateOnly),
		}); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error": "Unable to write CSV data",
			})
			return
		}
	}
}
