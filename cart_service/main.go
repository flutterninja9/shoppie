package main

import (
	"os"

	"github.com/flutterninja9/shoppie/cart_service/database"
	"github.com/flutterninja9/shoppie/cart_service/router"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func main() {
	envLoadErr := godotenv.Load(".env")

	if envLoadErr != nil {
		logger.Panic(envLoadErr)
	}
	database.InitDb(logger)

	app := fiber.New()
	router.SetupRoutes(app, logger)
	logger.Fatal(app.Listen(os.Getenv("SERVER_PORT")))
}
