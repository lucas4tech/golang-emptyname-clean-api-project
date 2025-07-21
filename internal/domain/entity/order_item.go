package entity

import (
	"time"

	"app-challenge/internal/domain/exception"
	"app-challenge/internal/domain/value_object"
)

type OrderItem struct {
	ID        *value_object.UUID
	ProductID string
	OrderID   string
	Quantity  int
	Price     *value_object.Money
	CreatedAt time.Time
}

func NewOrderItem(productID string, quantity int, price *value_object.Money) (*OrderItem, error) {
	if productID == "" {
		return nil, exception.NewRequiredFieldError("OrderItem", "productID")
	}
	if quantity <= 0 {
		return nil, exception.NewInvalidFieldError(quantity, "quantity must be positive")
	}
	if price == nil {
		return nil, exception.NewRequiredFieldError("OrderItem", "unitPrice")
	}

	now := time.Now()
	return &OrderItem{
		ID:        value_object.NewUUIDv4(),
		ProductID: productID,
		Quantity:  quantity,
		Price:     price,
		CreatedAt: now,
	}, nil
}
func (oi *OrderItem) UpdateQuantity(quantity int) error {
	if quantity <= 0 {
		return exception.NewInvalidFieldError(quantity, "quantity must be positive")
	}
	oi.Quantity = quantity
	return nil
}

func (oi *OrderItem) UpdatePrice(price *value_object.Money) error {
	if price == nil {
		return exception.NewRequiredFieldError("OrderItem", "price")
	}
	oi.Price = price
	return nil
}

func (oi *OrderItem) GetTotalPrice() (*value_object.Money, error) {
	return oi.Price.Multiply(float64(oi.Quantity))
}
