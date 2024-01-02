package cartsdk

import (
	"encoding/json"
	"io"

	"github.com/google/uuid"
)

type CartModel struct {
	ID     uuid.UUID  `json:"id"`
	UserID uuid.UUID  `json:"user_id"`
	Items  []CartItem `json:"items"`
}

func CartModelToJson(r io.Reader) (*CartModel, error) {
	cart := new(CartModel)
	err := json.NewDecoder(r).Decode(cart)

	if err != nil {
		return nil, err
	}

	return cart, nil
}
