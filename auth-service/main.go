package main

import (
	"auth-service/configs"
	"auth-service/grpc"
	"auth-service/routes"
	_ "embed"
	"log"
	"net"

	pb "auth-service/protos"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	rpc "google.golang.org/grpc"
)

//go:embed .env
var envData []byte

func main() {
	configs.LoadEnv(string(envData))

	configs.ConnectDB()

	ginServer := gin.Default()
	grpcServer := rpc.NewServer()

	ginServer.Use(cors.New(cors.Config{
		AllowOrigins:     []string{configs.GetEnv("FRONTEND_ADDRESS", "http://localhost:5173")},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	routes.SetupRoutes(ginServer)
	pb.RegisterAuthServiceServer(grpcServer, &grpc.AuthServer{})

	go func() {
		listen, err := net.Listen("tcp", ":50051")
		if err != nil {
			log.Fatalf("Unable to listen the port 50051 for gRPC server: %s\n", err)
		}

		err = grpcServer.Serve(listen)
		if err != nil {
			log.Fatalf("Unable to server the gRPC server: %s\n", err)
		}

		log.Println("Listening and serving gRPC on :50051")
	}()

	ginServer.Run(configs.GetEnv("ADDRESS", ":8080"))
}
