package middleware

import (
	"log"
	"net/http"
	"strings"
	"time"

	"context"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"employee-service/configs"
	pb "employee-service/protos"
)

func Check() gin.HandlerFunc {
	return func(c *gin.Context) {
		text := c.GetHeader("Authorization")

		grpcAddress := configs.GetEnv("GRPC_ADDRESS", "localhost:50051")

		connection, err := grpc.NewClient(grpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
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
			msg := err.Error()
			parts := strings.Split(msg, "desc =")
			if len(parts) > 1 {
				msg = strings.TrimSpace(parts[1])
				msg = cases.Title(language.Tag{}).String(msg)
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": msg,
			})
			return
		} else if response.Username == "" || response.EmployeeId == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "You are unauthorized and unable to proceed with the request.",
			})
			return
		}

		c.Set("username", response.Username)
		c.Set("employee_id", response.EmployeeId)

		log.Println("You are authorized and proceeding with the request.")
		c.Next()
	}
}
