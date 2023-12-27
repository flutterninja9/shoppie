package routers

import (
	"github.com/flutterninja9/shoppie/shoppie_bff/cart/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func SetupRouters(l *logrus.Logger, a *fiber.App) error {
	api := a.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/cart", func(c *fiber.Ctx) error {
		return controllers.GetCart(c, l)
	})

	v1.Post("/cart/items", func(c *fiber.Ctx) error {
		return controllers.AddToCart(c, l)
	})

	v1.Put("/cart/items/:itemId", func(c *fiber.Ctx) error {
		return controllers.UpdateCartItem(c, l)
	})

	v1.Delete("/cart/items/:itemId", func(c *fiber.Ctx) error {
		return controllers.RemoveFromCart(c, l)
	})

	l.Info("Cart router -> ✅")
	return nil
}
