package main

import (
	"os"

	"github.com/flutterninja9/shoppie/order_service/database"
	"github.com/flutterninja9/shoppie/order_service/router"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func main() {
	envErr := godotenv.Load(".env")

	if envErr != nil {
		logger.Fatal(envErr)
		panic(envErr)
	}

	dbErr := database.InitDb(logger)
	if dbErr != nil {
		logger.Fatal(dbErr)
		panic(dbErr)
	}

	app := fiber.New()
	router.SetupRoutes(app, logger)
	logger.Fatal(app.Listen(os.Getenv("SERVER_PORT")))
	logger.Info("Server stopped")
}
