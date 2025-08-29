package handlers

import (
	"encoding/csv"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DownloadTemplate(context *gin.Context) {
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
			"error": "Unable to write CSV header.",
		})
		return
	}
}
