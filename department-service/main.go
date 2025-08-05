package main

import (
	"department-service/configs"
	"department-service/routes"
	_ "embed"

	"github.com/gin-gonic/gin"
)

//go:embed .env
var envData []byte

func main() {
	configs.LoadEnv(string(envData))

	configs.ConnectDB()

	engine := gin.Default()

	routes.SetupRoutes(engine)

	engine.Run(configs.GetEnv("ADDRESS", ":8080"))
}
