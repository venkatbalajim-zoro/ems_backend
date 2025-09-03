package main

import (
	_ "embed"
	"employee-service/configs"
	"employee-service/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//go:embed .env
var envData []byte

func main() {
	configs.LoadEnv(string(envData))

	configs.ConnectDB()

	engine := gin.Default()
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{configs.GetEnv("FRONTEND_ADDRESS", "http://localhost:5173")},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	routes.SetupRoutes(engine)

	engine.Run(configs.GetEnv("ADDRESS", ":3000"))
}
