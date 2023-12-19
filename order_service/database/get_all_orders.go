package database

import (
	"order_service/models"

	"github.com/sirupsen/logrus"
)

func GetAllOrders(logger *logrus.Logger) ([]models.OrderModel, error) {
	logger.Info("Fetching all orders")

	var orders []models.OrderModel
	result := DbInstance.Find(&orders)

	if result.Error != nil {
		logger.Errorf("Failed to get orders %v", result.Error)
		return nil, result.Error
	}

	logger.Infof("Fetched %d total orders", len(orders))
	return orders, nil
}
