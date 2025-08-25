package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "You are logged in already",
	})
}
