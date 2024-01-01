package ordersdk

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/google/uuid"
)

type Order struct {
	ID          uuid.UUID   `json:"id"`
	Status      string      `json:"status"`
	TotalAmount float64     `json:"total_amount"`
	OrderItems  []OrderItem `json:"order_items"`
}

type Orders []Order

func (o *Order) ToJson() (io.Reader, error) {
	buffer := new(bytes.Buffer)
	err := json.NewEncoder(buffer).Encode(o)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func OrderFromJson(r io.Reader) (*Order, error) {
	order := new(Order)
	err := json.NewDecoder(r).Decode(order)
	if err != nil {
		return nil, err
	}

	return order, nil
}


func OrdersFromJson(r io.Reader) (*Orders, error) {
	orders := new(Orders)
	err := json.NewDecoder(r).Decode(orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
