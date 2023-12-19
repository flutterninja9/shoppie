package controllers

import (
	"github.com/flutterninja9/shoppie/order_service/database"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func GetAllOrders(c *fiber.Ctx, logger *logrus.Logger) error {
	orders, err := database.GetAllOrders(logger)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.Status(200).JSON(orders)
}
func CreateNewOrder(c *fiber.Ctx, logger *logrus.Logger) error {
	return c.SendString("CreateNewOrder")
}
func GetOrderById(c *fiber.Ctx, logger *logrus.Logger) error {
	orderId := c.Params("id")
	order, err := database.GetOrderById(orderId, logger)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.Status(200).JSON(order)
}
func GetOrdersOfUser(c *fiber.Ctx, logger *logrus.Logger) error {
	userId := c.Params("userId")
	orders, err := database.GetUserOrders(userId, logger)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.Status(200).JSON(orders)
}
func UpdateOrder(c *fiber.Ctx, logger *logrus.Logger) error {
	return c.SendString("UpdateOrder")
}
func CancelOrder(c *fiber.Ctx, logger *logrus.Logger) error {
	orderId := c.Params("id")
	err := database.CancelOrder(orderId, logger)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Cancelled",
	})
}
