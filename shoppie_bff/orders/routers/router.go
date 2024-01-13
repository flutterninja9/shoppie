package routers

import (
	middleware "github.com/flutterninja9/shoppie/shoppie_bff/middlewares"
	"github.com/flutterninja9/shoppie/shoppie_bff/orders/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

func SetupRouters(a *fiber.App, container *dig.Container) error {
	api := a.Group("/api")
	v1 := api.Group("/v1")
	orders := v1.Group("/orders")
	orders.Use(middleware.AuthMiddleware)

	orders.Get("/", func(c *fiber.Ctx) error {
		return controllers.GetUserOrders(c, container)
	})

	orders.Get("/all", func(c *fiber.Ctx) error {
		return controllers.GetAllOrders(c, container)
	})

	orders.Post("/", func(c *fiber.Ctx) error {
		return controllers.PlaceOrder(c, container)
	})

	orders.Get("/:id/cancel", func(c *fiber.Ctx) error {
		return controllers.CancelOrder(c, container)
	})

	container.Invoke(func(l *logrus.Logger) error {
		l.Info("Orders router -> ✅")
		return nil
	})

	return nil
}
