package main

import (
	"os"

	ordersdk "github.com/flutterninja9/shoppie/order_sdk"
	"github.com/flutterninja9/shoppie/user_service/database"
	"github.com/flutterninja9/shoppie/user_service/router"

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

	app := fiber.New()
	logger.Info("Server started")
	orderSdk := ordersdk.NewOrderSdk(os.Getenv("ORDER_SERVICE_BASE_URL"))
	router.SetupRoutes(app, logger, &orderSdk)
	logger.Fatal(app.Listen(os.Getenv("SERVER_PORT")))
	logger.Info("Server stopped")
}
