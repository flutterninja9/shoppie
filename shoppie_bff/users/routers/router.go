package routers

import (
	"github.com/flutterninja9/shoppie/shoppie_bff/users/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func SetupRouters(l *logrus.Logger, a *fiber.App) error {
	api := a.Group("/api")
	v1 := api.Group("/v1")

	v1.Post("/login", func(c *fiber.Ctx) error {
		return controllers.Login(c, l)
	})
	v1.Post("/register", func(c *fiber.Ctx) error {
		return controllers.Register(c, l)
	})
	v1.Get("/user", func(c *fiber.Ctx) error {
		return controllers.GetUser(c, l)
	})

	l.Info("User router -> ✅")
	return nil
}
