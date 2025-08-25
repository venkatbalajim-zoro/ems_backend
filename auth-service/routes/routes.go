package routes

import (
	"auth-service/handlers"
	"auth-service/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	routes := router.Group("/auth")
	{
		routes.POST("/login", handlers.Login)
		routes.POST("/recover", handlers.Recovery)
		routes.GET("/check", middleware.Check(), handlers.Ping)
		routes.POST("/register", middleware.Check(), handlers.Register)
		routes.DELETE("/remove", middleware.Check(), handlers.Remove)
		routes.GET("/accounts", middleware.Check(), handlers.Read)
		routes.DELETE("/delete", middleware.Check(), handlers.Delete)
	}
}
