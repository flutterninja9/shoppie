package database

import (
	"github.com/flutterninja9/shoppie/cart_service/models"
	"github.com/google/uuid"
)

func GetCartOrCreate(userId uuid.UUID) (*models.CartModel, error) {
	cart := new(models.CartModel)

	result := DbInstance.Preload("Items").FirstOrCreate(cart, models.CartModel{UserID: userId})

	if result.Error != nil {
		return nil, result.Error
	}

	return cart, nil
}
