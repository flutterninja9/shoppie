package main

import (
	"os"

	ordersdk "github.com/flutterninja9/shoppie/order_sdk"
	"github.com/flutterninja9/shoppie/user_service/database"
	"github.com/flutterninja9/shoppie/user_service/protos"
	"github.com/flutterninja9/shoppie/user_service/router"
	rpcserver "github.com/flutterninja9/shoppie/user_service/rpc_server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func main() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		logger.Fatal("Unable to load env file")
		panic(envErr)
	}

	dbErr := database.InitializeDb(logger)
	if dbErr != nil {
		panic(dbErr)
	}

	gs := grpc.NewServer()
	us := rpcserver.NewUserRPCServer()
	protos.RegisterRPCUserServiceServer(gs, us)
	reflection.Register(gs)
	// l, e := net.Listen("tcp", ":9090")
	// if e != nil {
	// 	logger.Error("Unable to listen")
	// 	panic(e)
	// }

	app := fiber.New()
	logger.Info("Server started")
	orderSdk := ordersdk.NewOrderSdk(os.Getenv("ORDER_SERVICE_BASE_URL"))
	router.SetupRoutes(app, logger, &orderSdk)
	logger.Fatal(app.Listen(os.Getenv("SERVER_PORT")))
	// logger.Fatal(gs.Serve(l))
	logger.Info("Server stopped")
}
