package main

import (
	"fmt"

	usersdk "github.com/flutterninja9/shoppie/user_sdk"
	"github.com/flutterninja9/shoppie/user_service/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println(err)
		return
	}

	client := protos.NewRPCUserServiceClient(conn)
	sdk := usersdk.NewUserService("http://localhost:3000/users/v1", client)

	// Health Check
	msg, err := sdk.HealthCheck()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(msg)
	// // Get user
	// user, _ := sdk.GetUserById("_userid_", "_token_")

	// fmt.Println(user)

	// // Login
	// loginRes, _ := sdk.Login(&usersdk.LoginRequest{Email: "test@gmail.com", Password: "test"})

	// fmt.Println(loginRes)

	// // Register
	// userRes, _ := sdk.Register(&usersdk.RegisterRequest{Email: "test1@gmail.com", Password: "test", FirstName: "Test", LastName: "Test", Username: "test1"})

	// fmt.Println(userRes)
}
