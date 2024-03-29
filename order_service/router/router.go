package router

import (
	"github.com/flutterninja9/shoppie/order_service/controllers"
	"github.com/flutterninja9/shoppie/order_service/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func SetupRoutes(app *fiber.App, logger *logrus.Logger) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(200)
	})

	orderPath := app.Group("/orders")
	orderPath.Use(middlewares.AuthMiddleware)
	v1 := orderPath.Group("/v1")

	// Get All orders
	v1.Get("/", func(c *fiber.Ctx) error {
		return controllers.GetAllOrders(c, logger)
	})

	// Create a new order
	v1.Post(
		"/",
		middlewares.UserExistenceCheckMiddleware,
		func(c *fiber.Ctx) error {
			return controllers.CreateNewOrder(c, logger)
		})

	// Get order details
	v1.Get("/:id", func(c *fiber.Ctx) error {
		return controllers.GetOrderById(c, logger)
	})

	// Get orders of a user
	v1.Get(
		"/user/:userId",
		middlewares.UserExistenceCheckMiddleware,
		func(c *fiber.Ctx) error {
			return controllers.GetOrdersOfUser(c, logger)
		})

	// Update order
	v1.Put("/:id",
		middlewares.UserExistenceCheckMiddleware,
		func(c *fiber.Ctx) error {
			return controllers.UpdateOrder(c, logger)
		})

	// Cancel order
	v1.Get(
		"/:id/cancel",
		middlewares.UserExistenceCheckMiddleware,
		func(c *fiber.Ctx) error {
			return controllers.CancelOrder(c, logger)
		})
}
