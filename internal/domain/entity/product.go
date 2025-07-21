package entity

import (
	"strings"
	"time"

	"app-challenge/internal/domain/exception"
	"app-challenge/internal/domain/value_object"
)

type Product struct {
	ID        *value_object.UUID
	Name      string
	Price     *value_object.Money
	Stock     int
	CreatedAt time.Time
}

func NewProduct(name string, price *value_object.Money, stock int) (*Product, error) {
	if name == "" {
		return nil, exception.NewRequiredFieldError("Product", "name")
	}

	if price == nil {
		return nil, exception.NewRequiredFieldError("Product", "price")
	}

	if stock < 0 {
		return nil, exception.NewInvalidFieldError(stock, "stock cannot be negative")
	}

	if len(strings.TrimSpace(name)) < 3 {
		return nil, exception.NewEntityValidationError("Product", "name", "name must have at least 3 characters")
	}

	if price.Amount() <= 0 {
		return nil, exception.NewEntityValidationError("Product", "price", "price must be positive")
	}

	now := time.Now()
	return &Product{
		ID:        value_object.NewUUIDv4(),
		Name:      strings.TrimSpace(name),
		Price:     price,
		Stock:     stock,
		CreatedAt: now,
	}, nil
}

func (p *Product) UpdateName(name string) error {
	if name == "" {
		return exception.NewRequiredFieldError("Product", "name")
	}

	if len(strings.TrimSpace(name)) < 3 {
		return exception.NewEntityValidationError("Product", "name", "name must have at least 3 characters")
	}

	p.Name = strings.TrimSpace(name)
	return nil
}

func (p *Product) UpdatePrice(price *value_object.Money) error {
	if price == nil {
		return exception.NewRequiredFieldError("Product", "price")
	}

	if price.Amount() <= 0 {
		return exception.NewEntityValidationError("Product", "price", "price must be positive")
	}

	p.Price = price
	return nil
}

func (p *Product) IncreaseStock(quantity int) error {
	if quantity <= 0 {
		return exception.NewInvalidFieldError(quantity, "quantity must be positive")
	}
	p.Stock += quantity
	return nil
}

func (p *Product) DecreaseStock(quantity int) error {
	if quantity <= 0 {
		return exception.NewInvalidFieldError(quantity, "quantity must be positive")
	}
	if p.Stock < quantity {
		return exception.NewEntityValidationError("Product", "stock", "insufficient stock")
	}
	p.Stock -= quantity
	return nil
}

func (p *Product) HasStock(quantity int) bool {
	return p.Stock >= quantity
}

func (p *Product) GetPriceInDollars() float64 {
	if p.Price == nil {
		return 0.0
	}
	return p.Price.Amount()
}

func (p *Product) IsAvailable() bool {
	return p.Stock > 0
}

func (p *Product) IsLowStock() bool {
	return p.Stock > 0 && p.Stock < 10
}

func (p *Product) IsOutOfStock() bool {
	return p.Stock == 0
}

func (p *Product) GetStockStatus() string {
	if p.IsOutOfStock() {
		return "Out of Stock"
	}
	if p.IsLowStock() {
		return "Low Stock"
	}
	return "In Stock"
}
