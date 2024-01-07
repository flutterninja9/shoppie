package main

import (
	cart "github.com/flutterninja9/shoppie/shoppie_bff/cart/routers"
	"github.com/flutterninja9/shoppie/shoppie_bff/config"
	"github.com/flutterninja9/shoppie/shoppie_bff/di"
	orders "github.com/flutterninja9/shoppie/shoppie_bff/orders/routers"
	products "github.com/flutterninja9/shoppie/shoppie_bff/products/routers"
	users "github.com/flutterninja9/shoppie/shoppie_bff/users/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

func main() {
	envErr := godotenv.Load(".env")

	if envErr != nil {
		panic(envErr)
	}

	// service locator
	sl := dig.New()
	di.SetupDi(sl)
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	users.SetupRouters(app, sl)
	products.SetupRouters(app, sl)
	orders.SetupRouters(app, sl)
	cart.SetupRouters(app, sl)

	sl.Invoke(func(l *logrus.Logger, c *config.AppConfig) error {
		l.Fatal(app.Listen(c.ServerPort))

		return nil
	})
}
