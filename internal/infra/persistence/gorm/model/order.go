package model

import (
	"time"
)

type Order struct {
	ID        string    `gorm:"type:uuid;primaryKey"`
	UserID    string    `gorm:"type:uuid;not null;index"`
	Total     float64   `gorm:"type:decimal(10,2);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`

	User User `gorm:"foreignKey:UserID;references:ID"`
}

func (Order) TableName() string {
	return "orders"
}

type OrderItem struct {
	ID        string    `gorm:"type:uuid;primaryKey"`
	OrderID   string    `gorm:"type:uuid;not null;index"`
	ProductID string    `gorm:"type:uuid;not null;index"`
	Quantity  int       `gorm:"not null"`
	Price     float64   `gorm:"type:decimal(10,2);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`

	Product Product `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Order   Order   `gorm:"foreignKey:OrderID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (OrderItem) TableName() string {
	return "order_items"
}
