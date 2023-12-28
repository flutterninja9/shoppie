package main

import (
	"fmt"

	usersdk "github.com/flutterninja9/shoppie/user_sdk"
)

func main() {
	sdk := usersdk.NewUserService("http://localhost:3000/users/v1")

	// Get user
	user, _ := sdk.GetUserById("_userid_", "_token_")

	fmt.Println(user)

	// Login
	loginRes, _ := sdk.Login(&usersdk.LoginRequest{Email: "test@gmail.com", Password: "test"})

	fmt.Println(loginRes)

	// Register
	userRes, _ := sdk.Register(&usersdk.RegisterRequest{Email: "test1@gmail.com", Password: "test", FirstName: "Test", LastName: "Test", Username: "test1"})

	fmt.Println(userRes)
}
