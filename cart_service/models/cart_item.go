package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartItem struct {
	ID          uuid.UUID `gorm:"primaryKey;not null;unique" json:"id"`
	CartID      uuid.UUID `gorm:"not null" json:"cart_id"`      // cart reference
	ProductID   uuid.UUID `gorm:"not null" json:"product_id"`   // product reference
	ProductName string    `gorm:"not null" json:"product_name"` // display name for the product
	Price       float64   `gorm:"not null" json:"price"`
	Quantity    int       `gorm:"default:1" json:"quantity"`
}

func (c *CartItem) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return nil
}
