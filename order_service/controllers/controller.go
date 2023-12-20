package controllers

import (
	"github.com/flutterninja9/shoppie/order_service/database"
	"github.com/flutterninja9/shoppie/order_service/middlewares"
	"github.com/flutterninja9/shoppie/order_service/models"
	"github.com/google/uuid"

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
	createOrderReq := new(CreateOrderRequest)
	parseErr := c.BodyParser(createOrderReq)
	if parseErr != nil {
		logger.Warning(parseErr)
		return c.SendStatus(fiber.ErrBadRequest.Code)
	}

	order := new(models.OrderModel)
	authInfo, ok := c.Locals("authInfo").(*middlewares.AuthInfo)

	if !ok {
		logger.Warning("No header found")
		return c.SendStatus(fiber.StatusBadRequest)
	}

	userId := authInfo.Claims["user_id"].(string)
	//TODO(AS): Also need to add a check for checking the passed products exists or not and add validation on their qty as well
	order.UserID, _ = uuid.Parse(userId)
	order.OrderItems = ToDBSaveableObject(createOrderReq.Items)
	var totalAmt float64

	for _, item := range order.OrderItems {
		totalAmt += float64(item.Quantity) * item.Price
	}
	order.TotalAmount = totalAmt

	createdOrder, creationErr := database.CreateOrder(logger, order)
	if creationErr != nil {
		logger.Warning(creationErr)
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.Status(200).JSON(createdOrder)
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
	authData, _ := c.Locals("authInfo").(*middlewares.AuthInfo)

	tokenId := authData.Claims["user_id"].(string)
	userId := c.Params("userId")

	// Check if token provided is of same user which is coming in param
	if tokenId != userId {
		logger.Warning("User authInfo mismatch")
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	orders, err := database.GetUserOrders(userId, logger)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.Status(200).JSON(orders)
}

func UpdateOrder(c *fiber.Ctx, logger *logrus.Logger) error {
	orderId := c.Params("id")

	order, err := database.GetOrderById(orderId, logger)

	if err != nil {
		logger.Warn(err)
		return c.SendStatus(404)
	}

	updateRequest := new(UpdateOrderStatusRequest)
	parseErr := c.BodyParser(updateRequest)

	if parseErr != nil {
		logger.Warn("Request body not correct")
		return c.SendStatus(fiber.ErrBadRequest.Code)
	}

	order.Status = string(updateRequest.Status)
	updateErr := database.UpdateOrder(&order, logger)

	if updateErr != nil {
		logger.Warn(updateErr)
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(200).JSON(order)
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
