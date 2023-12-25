package controllers

import (
	"github.com/flutterninja9/shoppie/cart_service/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func RemoveCartItem(c *fiber.Ctx, l *logrus.Logger) error {
	itemId := c.Params("itemId")
	itemUuid, _ := uuid.Parse(itemId)
	ok, err := database.DeleteItem(itemUuid)

	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Item removed successfully",
	})
}
