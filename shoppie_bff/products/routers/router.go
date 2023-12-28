package routers

import (
	"github.com/flutterninja9/shoppie/shoppie_bff/products/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

func SetupRouters(a *fiber.App, c *dig.Container) error {
	api := a.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/products", func(c *fiber.Ctx) error {
		return controllers.GetAllProducts(c)
	})
	v1.Get("/products/:productId", func(c *fiber.Ctx) error {
		return controllers.GetProductDetails(c)
	})

	c.Invoke(func(l *logrus.Logger) error {
		l.Info("Products router -> ✅")
		return nil
	})
	return nil
}
