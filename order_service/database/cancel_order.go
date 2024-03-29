package database

import (
	"github.com/flutterninja9/shoppie/order_service/globals"
	"github.com/sirupsen/logrus"
)

func CancelOrder(orderId string, logger *logrus.Logger) error {
	logger.Info("Cancelling order")

	order, fetchErr := GetOrderById(orderId, logger)
	if fetchErr != nil {
		logger.Errorf("Failed to get order %v", fetchErr)
		return fetchErr
	}

	order.Status = string(globals.Cancelled)
	result := DbInstance.Save(&order)

	if result.Error != nil {
		logger.Errorf("Failed to cancel order %v", result.Error)
		return result.Error
	}

	logger.Info("Cancelled order")
	return nil
}
