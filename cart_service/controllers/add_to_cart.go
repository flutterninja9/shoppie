package controllers

import (
	"github.com/flutterninja9/shoppie/cart_service/database"
	"github.com/flutterninja9/shoppie/cart_service/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func AddToCart(c *fiber.Ctx, l *logrus.Logger) error {
	userId := c.Params("userId")
	userUuid, uuidParseErr := uuid.Parse(userId)
	if uuidParseErr != nil {
		return c.SendStatus(fiber.ErrBadRequest.Code)
	}

	var items = []models.CartItem{}
	parseErr := c.BodyParser(&items)

	if parseErr != nil {
		return c.SendStatus(fiber.ErrBadRequest.Code)
	}

	addErr := database.AddItemToCart(userUuid, items)

	if addErr != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}
