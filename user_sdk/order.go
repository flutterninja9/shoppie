package usersdk

import (
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

func OrderFromJson(r io.Reader) *Order {
	order := new(Order)
	json.NewDecoder(r).Decode(order)
	return order
}

type Orders []Order

func OrdersFromJson(r io.Reader) *Orders {
	orders := new(Orders)
	json.NewDecoder(r).Decode(orders)
	return orders
}
