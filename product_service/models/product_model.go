package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductModel struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;" json:"id"`
	Name        string    `gorm:"type:varchar(100);not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Price       float64   `gorm:"not null" json:"price"`
	Quantity    int       `gorm:"not null" json:"quantity"`
	CreatedBy   string    `gorm:";not null" json:"created_by"`
}

func (order *ProductModel) BeforeCreate(tx *gorm.DB) (err error) {
	order.ID = uuid.New()
	return
}
