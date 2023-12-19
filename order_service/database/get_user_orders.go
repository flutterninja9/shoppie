package database

import (
	"order_service/models"

	"github.com/sirupsen/logrus"
)

func GetUserOrders(userId string, logger *logrus.Logger) ([]models.OrderModel, error) {
	logger.Info("Fetching user orders")

	var orders []models.OrderModel
	result := DbInstance.Find(&orders, "user_id=?", userId)

	if result.Error != nil {
		logger.Errorf("Failed to get orders %v", result.Error)
		return nil, result.Error
	}

	logger.Infof("Fetched %d orders", len(orders))
	return orders, nil
}
