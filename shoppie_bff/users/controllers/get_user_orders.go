package controllers

import (
	middleware "github.com/flutterninja9/shoppie/shoppie_bff/middlewares"
	usersdk "github.com/flutterninja9/shoppie/user_sdk"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

func GetUserOrders(c *fiber.Ctx, container *dig.Container) error {
	authInfo := c.Locals("authInfo").(*middleware.AuthInfo)

	return container.Invoke(func(l *logrus.Logger, s *usersdk.UserService) error {
		orders, err := s.GetOrders(authInfo.Token)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": orders,
		})
	})
}
