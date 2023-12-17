package main

import (
	"os"
	"user_service/database"
	"user_service/router"

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
	router.SetupRoutes(app, logger)
	logger.Fatal(app.Listen(os.Getenv("SERVER_PORT")))
	logger.Info("Server stopped")
}
