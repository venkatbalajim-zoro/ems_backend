package routes

import (
	"employee-service/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	routes := router.Group("/employees")
	{
		routes.GET("/read", handlers.Read)
		routes.POST("/add", handlers.Add)
		routes.PUT("/update", handlers.Update)
		routes.DELETE("/delete", handlers.Delete)
	}
}
