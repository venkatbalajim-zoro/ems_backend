package routes

import (
	"employee-service/handlers"
	"employee-service/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	routes := router.Group("/employees")
	{
		routes.GET("/read", middleware.Verify(), handlers.Read)
		routes.POST("/add", middleware.Verify(), handlers.Add)
		routes.POST("/upload-csv", middleware.Verify(), handlers.UploadCSV)
		routes.GET("/download-csv", middleware.Verify(), handlers.DownloadCSV)
		routes.PUT("/update", middleware.Verify(), handlers.Update)
		routes.DELETE("/delete", middleware.Verify(), handlers.Delete)
	}
}
