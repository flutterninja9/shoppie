package routers

import (
	"github.com/flutterninja9/shoppie/shoppie_bff/users/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

func SetupRouters(a *fiber.App, container *dig.Container) error {
	api := a.Group("/api")
	v1 := api.Group("/v1")

	v1.Post("/login", func(c *fiber.Ctx) error {
		return controllers.Login(c, container)
	})
	v1.Post("/register", func(c *fiber.Ctx) error {
		return controllers.Register(c, container)
	})
	v1.Get("/user", func(c *fiber.Ctx) error {
		return controllers.GetUser(c, container)
	})

	container.Invoke(func(l *logrus.Logger) error {
		l.Info("Users router -> ✅")
		return nil
	})
	return nil
}
