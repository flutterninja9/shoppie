package routers

import (
	"github.com/flutterninja9/shoppie/shoppie_bff/cart/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

func SetupRouters(a *fiber.App, container *dig.Container) error {
	api := a.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/cart", func(c *fiber.Ctx) error {
		return controllers.GetCart(c, container)
	})

	v1.Post("/cart/items", func(c *fiber.Ctx) error {
		return controllers.AddToCart(c, container)
	})

	v1.Put("/cart/items/:itemId", func(c *fiber.Ctx) error {
		return controllers.UpdateCartItem(c, container)
	})

	v1.Delete("/cart/items/:itemId", func(c *fiber.Ctx) error {
		return controllers.RemoveFromCart(c, container)
	})

	v1.Delete("/cart/clear", func(c *fiber.Ctx) error {
		return controllers.RemoveFromCart(c, container)
	})

	container.Invoke(func(l *logrus.Logger) error {
		l.Info("Cart router -> ✅")
		return nil
	})
	return nil
}
