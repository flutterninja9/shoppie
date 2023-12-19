package database

import (
	"github.com/flutterninja9/shoppie/order_service/models"

	"github.com/sirupsen/logrus"
)

func CreateOrder(logger *logrus.Logger, order *models.OrderModel) (*models.OrderModel, error) {
	logger.Info("Creating order")

	result := DbInstance.Create(order)

	if result.Error != nil {
		logger.Errorf("Failed to get orders %v", result.Error)
		return nil, result.Error
	}

	logger.Info("Order created sucessfully")
	return order, nil
}
