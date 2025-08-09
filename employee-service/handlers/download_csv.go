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
)

func DownloadCSV(context *gin.Context) {
	var input map[string]string
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to fetch the file path",
		})
		return
	}
	path := input["path"]

	file, err := os.Create(path)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to create a new file",
		})
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	var records []models.Employee
	configs.Database.Table("employees").Find(&records)

	writer.Write([]string{
		"Employee ID", "First Name", "Last Name", "Email ID", "Phone",
		"Gender", "Department ID", "Designation", "Salary", "Hire Date",
	})
	for _, record := range records {
		writer.Write([]string{
			strconv.Itoa(record.EmployeeID), record.FirstName, record.LastName,
			record.Email, record.Phone, record.Gender, strconv.Itoa(record.DepartmentID),
			strconv.FormatFloat(record.Salary, 'e', 2, 64), record.HireDate.Format(time.DateOnly),
		})
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Data is downloaded successfully",
	})
}
