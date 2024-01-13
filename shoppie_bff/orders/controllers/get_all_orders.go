package controllers

import (
	ordersdk "github.com/flutterninja9/shoppie/order_sdk"
	middleware "github.com/flutterninja9/shoppie/shoppie_bff/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

func GetAllOrders(c *fiber.Ctx, container *dig.Container) error {
	return container.Invoke(func(l *logrus.Logger, s ordersdk.OrderSdk) error {
		authInfo := c.Locals("authInfo").(*middleware.AuthInfo)
		orders, err := s.GetAllOrders(authInfo.Token)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": orders,
		})
	})
}
