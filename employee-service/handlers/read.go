package handlers

import (
	"employee-service/configs"
	"employee-service/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Read(context *gin.Context) {
	id, _ := context.Get("employee_id")
	var rows []models.Employee

	err := configs.Database.Table("employees").Find(&rows).Error
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error occured.",
		})
		return
	}

	for idx, row := range rows {
		deptId := row.DepartmentID
		data := make(map[string]any)
		err := configs.Database.Table("departments").Select("name").Where("id = ?", deptId).Scan(&data).Error
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error occured.",
			})
			return 
		} else {
			rows[idx].DepartmentName = data["name"].(string)
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "All employees details are fetched successfully.",
		"data":    rows,
		"employee_id": id,
	})
}
