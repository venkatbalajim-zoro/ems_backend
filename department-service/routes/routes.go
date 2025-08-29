package routes

import (
	"department-service/handlers"
	"department-service/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	routes := router.Group("/departments")
	{
		routes.GET("/read", middleware.Check(), handlers.Read)
		routes.POST("/add", middleware.Check(), handlers.Add)
		routes.PUT("/upload-csv", middleware.Check(), handlers.UploadCSV)
		routes.GET("/download-csv", middleware.Check(), handlers.DownloadCSV)
		routes.PUT("/update", middleware.Check(), handlers.Update)
		routes.DELETE("/delete", middleware.Check(), handlers.Delete)
		routes.GET("/download-template", middleware.Check(), handlers.DownloadTemplate)
	}
}
