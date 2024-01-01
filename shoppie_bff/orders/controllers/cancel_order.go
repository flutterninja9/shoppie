package controllers

import (
	ordersdk "github.com/flutterninja9/shoppie/order_sdk"
	middleware "github.com/flutterninja9/shoppie/shoppie_bff/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

func CancelOrder(c *fiber.Ctx, container *dig.Container) error {
	authInfo := c.Locals("authInfo").(*middleware.AuthInfo)
	oId := c.Params("id")
	return container.Invoke(func(l *logrus.Logger, s ordersdk.OrderSdk) error {
		res, err := s.CancelOrder(oId, authInfo.Token)

		if err != nil {
			l.Error(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": res,
		})
	})
}
