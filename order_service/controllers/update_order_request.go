package controllers

import "github.com/flutterninja9/shoppie/order_service/globals"

type UpdateOrderStatusRequest struct {
	Status globals.OrderStatus `json:"status"`
}
