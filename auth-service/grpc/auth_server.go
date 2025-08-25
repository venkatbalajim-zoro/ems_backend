package grpc

import (
	"context"

	pb "auth-service/protos"
	"auth-service/utils"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
}

func (server *AuthServer) VerifyToken(ctx context.Context, request *pb.VerifyTokenRequest) (*pb.VerifyTokenResponse, error) {
	token := request.GetText()

	username, employee_id, err := utils.Verify(token)

	if err != nil {
		return nil, err
	}

	return &pb.VerifyTokenResponse{
		Username:   username,
		EmployeeId: int64(employee_id),
	}, nil
}
