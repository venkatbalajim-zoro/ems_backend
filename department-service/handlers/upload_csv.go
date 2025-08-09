package handlers

import (
	"department-service/configs"
	"department-service/models"
	"encoding/csv"
	"errors"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
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
			"error": "Unable to read the CSV file",
		})
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to read the CSV content",
		})
		return
	}

	if len(records) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Emply CSV file",
		})
		return
	}

	var data []models.Department
	for _, record := range records {
		if len(record) < 2 || len(record) > 2 {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "Required two data - id and name",
			})
			return
		}
		id, err := strconv.Atoi(record[0])
		name := record[1]
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid data in the CSV file",
			})
			return
		}
		elt := models.Department{
			ID:   id,
			Name: name,
		}
		data = append(data, elt)
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
					"error": "Unable to add data as duplicate data exists",
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
		"message": "Data is uploaded successfully",
	})
}
