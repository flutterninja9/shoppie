package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderModel struct {
	gorm.Model
	ID          uuid.UUID   `gorm:"type:uuid;primaryKey;" json:"id"`
	UserID      uuid.UUID   `gorm:"type:uuid;not null" json:"user_id"`
	Status      string      `gorm:"type:varchar(20);not null" json:"status"`
	TotalAmount float64     `gorm:"not null" json:"total_amount"`
	OrderItems  []OrderItem `gorm:"foreignKey:OrderID" json:"order_items"`
}

func (order *OrderModel) BeforeCreate(tx *gorm.DB) (err error) {
	order.ID = uuid.New()
	return
}