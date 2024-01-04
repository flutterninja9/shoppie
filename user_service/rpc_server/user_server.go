package rpcserver

import (
	"context"

	"github.com/flutterninja9/shoppie/user_service_rpc"
)

type UserRPCServer struct {
	user_service_rpc.UnimplementedRPCUserServiceServer
}

func NewUserRPCServer() *UserRPCServer {
	return &UserRPCServer{}
}

func (UserRPCServer) HealthCheck(
	context.Context,
	*user_service_rpc.HealthCheckRequest,
) (*user_service_rpc.HealthResponse, error) {
	return &user_service_rpc.HealthResponse{
		Message: "OK",
	}, nil
}
