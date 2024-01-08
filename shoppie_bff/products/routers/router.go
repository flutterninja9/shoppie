package routers

import (
	middleware "github.com/flutterninja9/shoppie/shoppie_bff/middlewares"
	"github.com/flutterninja9/shoppie/shoppie_bff/products/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

func SetupRouters(a *fiber.App, container *dig.Container) error {
	api := a.Group("/api")
	v1 := api.Group("/v1")
	products := v1.Group("/products")
	products.Use(middleware.AuthMiddleware)

	products.Get("/", func(c *fiber.Ctx) error {
		return controllers.GetAllProducts(c, container)
	})

	products.Post("/", func(c *fiber.Ctx) error {
		return controllers.CreateProduct(c, container)
	})

	products.Get("/:productId", func(c *fiber.Ctx) error {
		return controllers.GetProductDetails(c, container)
	})

	products.Patch("/:productId", func(c *fiber.Ctx) error {
		return controllers.UpdateProduct(c, container)
	})

	products.Delete("/:productId", func(c *fiber.Ctx) error {
		return controllers.DeleteProduct(c, container)
	})

	container.Invoke(func(l *logrus.Logger) error {
		l.Info("Products router -> ✅")
		return nil
	})
	return nil
}
