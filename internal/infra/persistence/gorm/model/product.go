package model

import (
	"time"
)

type Product struct {
	ID        string    `gorm:"type:uuid;primaryKey"`
	Name      string    `gorm:"type:varchar(100);not null"`
	Price     float64   `gorm:"type:decimal(10,2);not null"`
	Stock     int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (Product) TableName() string {
	return "products"
}
