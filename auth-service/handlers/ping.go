package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(context *gin.Context) {
	username, _ := context.Get("username")
	employeeId, _ := context.Get("employee_id")

	context.JSON(http.StatusOK, gin.H{
		"message": "You are logged in already",
		"data": gin.H{
			"username": username,
			"employee_id": employeeId,
		},
	})
}