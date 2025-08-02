package routes

import (
	"auth-service/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	routes := router.Group("/auth")
	{
		routes.GET("/login", handlers.Login)
		routes.POST("/register", handlers.Register)
		routes.GET("/recover", handlers.Recovery)
	}
}
