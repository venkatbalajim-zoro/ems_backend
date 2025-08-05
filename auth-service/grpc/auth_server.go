package grpc

import (
	"context"

	"auth-service/middleware"
	pb "shared/protos"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
}

func (server *AuthServer) VerifyToken(ctx context.Context, request *pb.VerifyTokenRequest) (*pb.VerifyTokenResponse, error) {
	token := request.GetText()

	status := middleware.Verify(token)

	return &pb.VerifyTokenResponse{
		Response: status,
	}, nil
}
