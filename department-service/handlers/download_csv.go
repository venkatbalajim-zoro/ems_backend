package handlers

import (
	"department-service/configs"
	"department-service/models"
	"encoding/csv"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DownloadCSV(context *gin.Context) {
	var data []models.Department
	result := configs.Database.Table("departments").Find(&data)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to fetch data from database.",
		})
		return
	}

	context.Header("Content-Description", "File Transfer")
	context.Header("Content-Disposition", `attachment; filename="departments.csv"`)
	context.Header("Content-Type", "text/csv")

	writer := csv.NewWriter(context.Writer)
	defer writer.Flush()

	if err := writer.Write([]string{"ID", "Name"}); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to write CSV header.",
		})
		return
	}

	for _, record := range data {
		if err := writer.Write([]string{strconv.Itoa(record.ID), record.Name}); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error": "Unable to write CSV data.",
			})
			return
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Departments data is downloaded successfully.",
	})
}
