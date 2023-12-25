package controllers

import (
	"github.com/flutterninja9/shoppie/cart_service/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func GetCart(c *fiber.Ctx, l *logrus.Logger) error {
	userId := c.Params("userId")
	userUuid, parseErr := uuid.Parse(userId)

	if parseErr != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"error": "Not a valid userId",
		})
	}

	cart, err := database.GetCartOrCreate(userUuid)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(cart)
}
