package handlers

import (
	"department-service/configs"
	"department-service/models"
	"encoding/csv"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DownloadCSV(context *gin.Context) {
	var input map[string]string
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to fetch the input",
		})
	}
	path := input["path"]

	var data []models.Department
	configs.Database.Table("departments").Find(&data)

	file, err := os.Create(path)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to create a new file for downloading data",
		})
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"ID", "Name"})
	for _, record := range data {
		err := writer.Write([]string{strconv.Itoa(record.ID), record.Name})
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error": "Unable to save the data in the file",
			})
			writer.Flush()
			os.Remove(path)
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Data is downloaded successfully",
	})
}
