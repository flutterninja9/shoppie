package main

import (
	"log"
	"os"

	"github.com/flutterninja9/shoppie/log_service/db"
	"github.com/flutterninja9/shoppie/log_service/router"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	dbErr := db.SetupDb()
	if dbErr != nil {
		panic(err)
	}

	app := fiber.New()
	router.SetupRouter(app)
	log.Fatal(app.Listen(os.Getenv("PORT")))
}
