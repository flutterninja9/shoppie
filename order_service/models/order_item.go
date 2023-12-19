package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderItem struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;" json:"id"`
	OrderID   uuid.UUID `gorm:"type:uuid;not null" json:"order_id"`   // Reference to the Order
	ProductID uuid.UUID `gorm:"type:uuid;not null" json:"product_id"` // Reference to the Product
	Quantity  int       `gorm:"not null" json:"quantity"`
	Price     float64   `gorm:"not null" json:"price"` // Price at the time of order
}

func (item *OrderItem) BeforeCreate(tx *gorm.DB) (err error) {
	item.ID = uuid.New()
	return
}
