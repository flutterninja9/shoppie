package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartModel struct {
	ID     uuid.UUID  `gorm:"primaryKey;not null" json:"id"`
	UserID uuid.UUID  `gorm:"not null" json:"user_id"`
	Items  []CartItem `gorm:"foreignKey:CartID;default:[]" json:"items"`
}

func (c *CartModel) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return nil
}
