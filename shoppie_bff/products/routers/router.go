package routers

import (
	"github.com/flutterninja9/shoppie/shoppie_bff/products/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func SetupRouters(l *logrus.Logger, a *fiber.App) error {
	api := a.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/products", func(c *fiber.Ctx) error {
		return controllers.GetAllProducts(c, l)
	})
	v1.Get("/products/:productId", func(c *fiber.Ctx) error {
		return controllers.GetProductDetails(c, l)
	})

	l.Info("Product router -> ✅")
	return nil
}
