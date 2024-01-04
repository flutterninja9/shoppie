package rpcserver

import (
	"context"

	"github.com/flutterninja9/shoppie/user_service/protos"
)

type UserRPCServer struct {
	protos.UnimplementedRPCUserServiceServer
}

func NewUserRPCServer() *UserRPCServer {
	return &UserRPCServer{}
}

func (UserRPCServer) HealthCheck(
	context.Context,
	*protos.HealthCheckRequest,
) (*protos.HealthResponse, error) {
	return &protos.HealthResponse{
		Message: "OK",
	}, nil
}
