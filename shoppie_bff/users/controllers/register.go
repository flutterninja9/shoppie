package controllers

import (
	usersdk "github.com/flutterninja9/shoppie/user_sdk"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

func Register(c *fiber.Ctx, container *dig.Container) error {
	return container.Invoke(func(s *usersdk.UserService, l *logrus.Logger) error {
		l.Info("Parsing user request")
		regRequest := new(usersdk.RegisterRequest)
		parseErr := c.BodyParser(regRequest)

		if parseErr != nil {
			return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
				"error": parseErr,
			})
		}
		l.Info("Parsed user request", regRequest)
		l.Info("Calling user-service")
		tokenResponse, err := s.Register(regRequest)
		if err != nil {
			l.Fatal(err)
			return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
				"error": err,
			})
		}
		l.Info("Got response from user-service", tokenResponse)
		return c.Status(200).JSON(tokenResponse)
	})
}
