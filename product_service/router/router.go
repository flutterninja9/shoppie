package router

import (
	"product_service/controllers"
	"product_service/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func SetupRoutes(app *fiber.App, logger *logrus.Logger) {
	logger.Info("Setting up routes")
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status": "OK",
		})
	})

	products := app.Group("/products")
	v1 := products.Group("/v1")
	v1.Use(middlewares.AuthMiddleware)
	v1.Use(middlewares.UserCheckerMiddleWare)
	v1.Get("/", controllers.GetAllProducts)
	v1.Post("/", controllers.CreateProduct)
	v1.Get("/:id", controllers.GetProductById)
	v1.Put("/:id", controllers.UpdateProduct)
	v1.Delete("/:id", controllers.DeleteProduct)
	logger.Info("Routes setup done")
}
