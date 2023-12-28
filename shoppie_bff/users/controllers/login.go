package controllers

import (
	usersdk "github.com/flutterninja9/shoppie/user_sdk"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

func Login(c *fiber.Ctx, container *dig.Container) error {
	loginRequest := new(usersdk.LoginRequest)
	parseErr := c.BodyParser(loginRequest)

	if parseErr != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"error": parseErr,
		})
	}

	return container.Invoke(func(s *usersdk.UserService, l *logrus.Logger) error {
		tokenResponse, err := s.Login(loginRequest)
		if err != nil {
			l.Fatal(err)
			return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
				"error": err,
			})
		}

		return c.Status(200).JSON(tokenResponse)
	})
}
