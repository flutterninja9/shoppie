package database

import (
	"github.com/flutterninja9/shoppie/cart_service/models"
	"github.com/google/uuid"
)

func DeleteItem(itemId uuid.UUID) (bool, error) {
	item := new(models.CartItem)
	result := DbInstance.First(item, "id=?", itemId)

	if result.Error != nil {
		return false, result.Error
	}

	deletedResult := DbInstance.Delete(item)
	if deletedResult.Error != nil {
		return false, deletedResult.Error
	}

	return true, nil
}
