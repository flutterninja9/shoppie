package router

import (
	ordersdk "github.com/flutterninja9/shoppie/order_sdk"
	"github.com/flutterninja9/shoppie/user_service/controllers"
	"github.com/flutterninja9/shoppie/user_service/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func SetupRoutes(app *fiber.App, logger *logrus.Logger, orderSdk *ordersdk.OrderSdk) {
	logger.Info("Setting up routes")
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "OK",
		})
	})

	userApp := app.Group("/users")
	v1 := userApp.Group("/v1")

	v1.Post("/register", func(c *fiber.Ctx) error {
		return controllers.Register(c, logger)
	})

	v1.Post("/login", func(c *fiber.Ctx) error {
		return controllers.Login(c, logger)
	})

	v1.Get("/user", middleware.AuthMiddleware, func(c *fiber.Ctx) error {
		return controllers.GetUserById(c, logger)
	})

	v1.Get("/orders", middleware.AuthMiddleware, func(c *fiber.Ctx) error {
		return controllers.GetUserOrders(c, logger, orderSdk)
	})

	logger.Info("Routes setup sucessfully")
}
