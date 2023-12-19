package controllers

import (
	"github.com/flutterninja9/shoppie/order_service/models"
	"github.com/google/uuid"
)

type CreateOrderRequest struct {
	Items []OrderItemRequest `json:"items"`
}

func ToDBSaveableObject(rec []OrderItemRequest) []models.OrderItem {
	var dbObjects []models.OrderItem

	for _, item := range rec {
		dbItem := models.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     item.Price,
		}

		dbObjects = append(dbObjects, dbItem)
	}

	return dbObjects
}

type OrderItemRequest struct {
	ProductID uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Price     float64   `json:"price"`
}
