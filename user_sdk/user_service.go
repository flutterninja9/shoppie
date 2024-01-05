package usersdk

import (
	protos "github.com/flutterninja9/shoppie/user_service/protos"
)

type UserService struct {
	b string
	c protos.RPCUserServiceClient
}

func NewUserService(b string, c protos.RPCUserServiceClient) *UserService {
	return &UserService{
		b: b,
		c: c,
	}
}
