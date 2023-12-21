package ordersdk

import "github.com/google/uuid"

type Order struct {
	ID          uuid.UUID   `json:"id"`
	Status      string      `json:"status"`
	TotalAmount float64     `json:"total_amount"`
	OrderItems  []OrderItem `json:"order_items"`
}
