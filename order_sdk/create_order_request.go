package ordersdk

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/google/uuid"
)

type CreateOrderRequest struct {
	Items []OrderItemRequest `json:"items"`
}

func (c *CreateOrderRequest) ToJson() (io.Reader, error) {
	buffer := new(bytes.Buffer)
	err := json.NewEncoder(buffer).Encode(c)

	if err != nil {
		return nil, err
	}
	return buffer, nil
}

type OrderItemRequest struct {
	ProductID uuid.UUID `json:"product_id" validate:"required"`
	Quantity  int       `json:"quantity" validate:"gt=0"`
	Price     float64   `json:"price" validate:"required"`
}
