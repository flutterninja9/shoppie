package router

import (
	"github.com/flutterninja9/shoppie/cart_service/controllers"
	"github.com/flutterninja9/shoppie/cart_service/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func SetupRoutes(app *fiber.App, logger *logrus.Logger) {
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("OK")
	})

	cart := app.Group("/cart")
	v1 := cart.Group("/v1")
	v1.Use(middlewares.AuthMiddleware)
	v1.Use(middlewares.UserExistenceCheckMiddleware)

	// Get user's cart
	v1.Get(":userId", func(c *fiber.Ctx) error {
		return controllers.GetCart(c, logger)
	})

	// Add items to the cart
	v1.Post(":userId/items", func(c *fiber.Ctx) error {
		return controllers.AddToCart(c, logger)
	})

	// Update any item already present in the cart
	v1.Put(":userId/items/:itemId", func(c *fiber.Ctx) error {
		return controllers.UpdateCartItem(c, logger)
	})

	// Remove any item present in the cart
	v1.Delete(":userId/items/:itemId", func(c *fiber.Ctx) error {
		return controllers.RemoveCartItem(c, logger)
	})

	// Remove all items present in the cart
	v1.Delete(":userId", func(c *fiber.Ctx) error {
		return controllers.ClearCart(c, logger)
	})
}
