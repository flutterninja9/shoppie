package controllers

import (
	ordersdk "github.com/flutterninja9/shoppie/order_sdk"
	middleware "github.com/flutterninja9/shoppie/shoppie_bff/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

func PlaceOrder(c *fiber.Ctx, container *dig.Container) error {
	authInfo := c.Locals("authInfo").(*middleware.AuthInfo)
	orderRequest := new(ordersdk.CreateOrderRequest)
	err := c.BodyParser(orderRequest)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	return container.Invoke(func(l *logrus.Logger, s ordersdk.OrderSdk) error {
		res, err := s.CreateOrder(orderRequest, authInfo.Token)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": res,
		})
	})
}
