package database

import (
	"github.com/flutterninja9/shoppie/order_service/models"

	"github.com/sirupsen/logrus"
)

func GetOrderById(orderId string, logger *logrus.Logger) (models.OrderModel, error) {
	logger.Info("Fetching user orders")

	order := new(models.OrderModel)
	result := DbInstance.First(&order, "id=?", orderId)

	if result.Error != nil {
		logger.Errorf("Failed to get orders %v", result.Error)
		return *order, result.Error
	}

	logger.Info("Fetched order")
	return *order, nil
}
