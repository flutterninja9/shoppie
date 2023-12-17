package router

import (
	"user_service/controllers"
	"user_service/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func SetupRoutes(app *fiber.App, logger *logrus.Logger) {
	logger.Info("Setting up routes")
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "OK",
		})
	})

	userApp := app.Group("/users")
	v1 := userApp.Group("/v1")

	v1.Post("/register", func(c *fiber.Ctx) error {
		return controllers.Register(c, logger)
	})
	v1.Post("/login", func(c *fiber.Ctx) error {
		return controllers.Login(c, logger)
	})
	v1.Get("/user", middleware.AuthMiddleware, func(c *fiber.Ctx) error {
		return controllers.GetUserById(c, logger)
	})
	logger.Info("Routes setup sucessfully")
}
