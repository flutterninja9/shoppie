package cartsdk

import (
	"encoding/json"
	"io"

	"github.com/google/uuid"
)

type CartItem struct {
	ID          uuid.UUID `json:"id"`
	CartID      uuid.UUID `json:"cart_id"`      // cart reference
	ProductID   uuid.UUID `json:"product_id"`   // product reference
	ProductName string    `json:"product_name"` // display name for the product
	Price       float64   `json:"price"`
	Quantity    int       `json:"quantity"`
}

func CartItemToJson(r io.Reader) (*CartItem, error) {
	item := new(CartItem)
	err := json.NewDecoder(r).Decode(item)

	if err != nil {
		return nil, err
	}

	return item, nil
}
