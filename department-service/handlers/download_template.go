package handlers

import (
	"encoding/csv"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DownloadTemplate(context *gin.Context) {
	context.Header("Content-Description", "File Transfer")
	context.Header("Content-Disposition", `attachment; filename="departments-template.csv"`)
	context.Header("Content-Type", "text/csv")

	writer := csv.NewWriter(context.Writer)
	defer writer.Flush()

	if err := writer.Write([]string{"ID", "Name"}); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to write CSV header.",
		})
		return
	}
}
