package database

import (
	"github.com/flutterninja9/shoppie/order_service/models"
	"github.com/sirupsen/logrus"
)

func UpdateOrder(order *models.OrderModel, logger *logrus.Logger) error {
	logger.Info("Updating order")

	result := DbInstance.Save(&order)

	if result.Error != nil {
		logger.Errorf("Failed to update order %v", result.Error)
		return result.Error
	}

	logger.Info("Order updated")
	return nil
}
