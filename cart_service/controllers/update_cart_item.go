package controllers

import (
	"github.com/flutterninja9/shoppie/cart_service/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type UpdateCartItemRequest struct {
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

func UpdateCartItem(c *fiber.Ctx, l *logrus.Logger) error {
	itemId := c.Params("itemId")
	itemUuid, _ := uuid.Parse(itemId)
	cartItem, err := database.GetItem(itemUuid)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}

	request := new(UpdateCartItemRequest)
	parseErr := c.BodyParser(request)
	if parseErr != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": parseErr,
		})
	}

	cartItem.Price = request.Price
	cartItem.Quantity = request.Quantity

	ok, updateErr := database.UpdateItem(cartItem)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": updateErr,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Item updated successfully",
	})
}
