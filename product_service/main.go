package main

import (
	"os"
	"product_service/database"
	"product_service/router"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func main() {
	logger.Info("Starting server")
	envErr := godotenv.Load(".env")
	if envErr != nil {
		logger.Fatal("Unable to load env")
		panic(envErr)
	}
	logger.Info("Env variables loaded")
	dbErr := database.InitDb(logger)

	if dbErr != nil {
		logger.Fatal("Unable to initialize db")
		panic(dbErr)
	}
	logger.Info("Initializing app")
	app := fiber.New()
	logger.Info("App initialized")
	router.SetupRoutes(app, logger)
	logger.Fatal(app.Listen(os.Getenv("SERVER_PORT")))
	logger.Info("App stopped")
}
