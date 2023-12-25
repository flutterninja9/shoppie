package database

import (
	"github.com/flutterninja9/shoppie/cart_service/models"
)

func UpdateItem(item *models.CartItem) (bool, error) {
	result := DbInstance.Save(item)

	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}
