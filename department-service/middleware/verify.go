package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"

	context "context"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"department-service/configs"
	pb "department-service/protos"
)

func Verify() gin.HandlerFunc {
	return func(c *gin.Context) {
		text := c.GetHeader("Authorization")

		grpcAddress := configs.GetEnv("GRPC_ADDRESS", "localhost:50051")

		connection, err := grpc.NewClient(grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Printf("Error in connecting gRPC: %s\n", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Unable to connect with the gRPC server.",
			})
			return
		}

		client := pb.NewAuthServiceClient(connection)

		context, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		response, err := client.VerifyToken(context, &pb.VerifyTokenRequest{
			Text: text,
		})

		if err != nil {
			errMessage := fmt.Sprintf("Unable to verify the token - %s", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": errMessage,
			})
			return
		} else if !response.Response {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "You are unauthorized and unable to proceed with the request.",
			})
			return
		}

		log.Println("You are authorized and proceeding with the request ...")
		c.Next()
	}
}
