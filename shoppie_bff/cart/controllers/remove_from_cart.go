package controllers

import (
	cartsdk "github.com/flutterninja9/shoppie/cart_sdk"
	middleware "github.com/flutterninja9/shoppie/shoppie_bff/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

// Deletes item from the cart
func RemoveFromCart(c *fiber.Ctx, container *dig.Container) error {
	authInfo := c.Locals("authInfo").(*middleware.AuthInfo)
	userId := authInfo.Claims["user_id"].(string)
	itemId := c.Params("itemId")

	return container.Invoke(func(s *cartsdk.CartSdk, l *logrus.Logger) error {
		response, err := s.RemoveCartItem(userId, itemId, authInfo.Token)

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err,
			})
		}

		return c.Status(fiber.StatusOK).JSON(response)
	})
}
