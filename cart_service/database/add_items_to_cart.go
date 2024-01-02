package database

import (
	"github.com/flutterninja9/shoppie/cart_service/models"
	"github.com/google/uuid"
)

func AddItemToCart(userId uuid.UUID, items []models.CartItem) error {
	cart := new(models.CartModel)

	result := DbInstance.Preload("Items").FirstOrCreate(cart, "user_id=?", userId)

	if result.Error != nil {
		return result.Error
	}

	cart.Items = append(cart.Items, items...)
	updateResult := DbInstance.Save(cart)

	if updateResult.Error != nil {
		return updateResult.Error
	}

	return nil
}
