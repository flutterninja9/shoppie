package database

import (
	"github.com/flutterninja9/shoppie/cart_service/models"
	"github.com/google/uuid"
)

func GetItem(itemId uuid.UUID) (*models.CartItem, error) {
	item := new(models.CartItem)

	result := DbInstance.First(item, "id=?", itemId)

	if result.Error != nil {
		return nil, result.Error
	}

	return item, nil
}
