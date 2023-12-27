package main

import (
	cart "github.com/flutterninja9/shoppie/shoppie_bff/cart/routers"
	"github.com/flutterninja9/shoppie/shoppie_bff/config"
	orders "github.com/flutterninja9/shoppie/shoppie_bff/orders/routers"
	products "github.com/flutterninja9/shoppie/shoppie_bff/products/routers"
	users "github.com/flutterninja9/shoppie/shoppie_bff/users/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()
var Config *config.AppConfig

func main() {
	envErr := godotenv.Load(".env")

	if envErr != nil {
		logger.Panic(envErr)
	}

	Config = config.NewAppConfig()

	if Config == nil {
		logger.Fatal("Unable to load app config")
	}

	app := fiber.New()

	users.SetupRouters(logger, app)
	products.SetupRouters(logger, app)
	orders.SetupRouters(logger, app)
	cart.SetupRouters(logger, app)

	logger.Fatal(app.Listen(Config.ServerPort))
}
