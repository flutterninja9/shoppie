package database

import (
	"github.com/flutterninja9/shoppie/cart_service/models"
	"github.com/google/uuid"
)

func ClearCart(userId uuid.UUID) (bool, error) {
	cart := new(models.CartModel)

	result := DbInstance.First(cart, "user_id=?", userId)

	if result.Error != nil {
		return false, result.Error
	}

	err := DbInstance.Where("cart_id=?", cart.ID).Delete(&models.CartItem{}).Error

	if err != nil {
		return false, err
	}
	return true, nil
}
