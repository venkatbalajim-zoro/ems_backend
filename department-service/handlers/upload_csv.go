package handlers

import (
	"department-service/configs"
	"department-service/models"
	"encoding/csv"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm/clause"
)

func UploadCSV(context *gin.Context) {
	file, err := context.FormFile("file")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "CSV file is required",
		})
		return
	}

	src, err := file.Open()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to open the uploaded file",
		})
		return
	}
	defer src.Close()

	reader := csv.NewReader(src)
	records, err := reader.ReadAll()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to read the CSV content",
		})
		return
	}

	if len(records) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Empty CSV file",
		})
		return
	}

	var data []models.Department
	for _, record := range records {
		if len(record) != 2 {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "Each row must have exactly 2 values: id and name",
			})
			return
		}

		id, err := strconv.Atoi(record[0])
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid ID in CSV",
			})
			return
		}

		data = append(data, models.Department{
			ID:   id,
			Name: record[1],
		})
	}

	err = configs.Database.Table("departments").Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"name"}),
	}).Create(&data).Error

	if err != nil {
		var sqlError *mysql.MySQLError
		if errors.As(err, &sqlError) {
			switch sqlError.Number {
			case 1062:
				context.JSON(http.StatusConflict, gin.H{
					"error": "Duplicate data exists",
				})
				return
			default:
				context.JSON(http.StatusInternalServerError, gin.H{
					"error": "Unable to add the details",
				})
				return
			}
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Data uploaded successfully",
	})
}
