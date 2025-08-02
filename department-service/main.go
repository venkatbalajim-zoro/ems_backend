package main

import (
	"department-service/configs"
	"department-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.LoadEnv()

	configs.ConnectDB()

	engine := gin.Default()

	routes.SetupRoutes(engine)

	engine.Run(configs.GetEnv("ADDRESS", ":8080"))
}
