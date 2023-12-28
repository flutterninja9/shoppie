package routers

import (
	"github.com/flutterninja9/shoppie/shoppie_bff/orders/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

func SetupRouters(a *fiber.App, c *dig.Container) error {
	api := a.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/orders", func(c *fiber.Ctx) error {
		return controllers.GetUserOrders(c)
	})

	v1.Post("/orders", func(c *fiber.Ctx) error {
		return controllers.PlaceOrder(c)
	})
	c.Invoke(func(l *logrus.Logger) error {
		l.Info("Orders router -> ✅")
		return nil
	})

	return nil
}
