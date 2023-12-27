package routers

import (
	"github.com/flutterninja9/shoppie/shoppie_bff/orders/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func SetupRouters(l *logrus.Logger, a *fiber.App) error {
	api := a.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/orders", func(c *fiber.Ctx) error {
		return controllers.GetUserOrders(c, l)
	})

	v1.Post("/orders", func(c *fiber.Ctx) error {
		return controllers.PlaceOrder(c, l)
	})
	l.Info("Orders router -> ✅")
	return nil
}
