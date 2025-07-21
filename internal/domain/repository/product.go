package repository

import (
	"app-challenge/internal/domain/entity"
	"context"
)

type ProductRepository interface {
	Create(ctx context.Context, product *entity.Product) error
	FindByID(ctx context.Context, id string) (*entity.Product, error)
	Update(ctx context.Context, product *entity.Product) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context, limit, offset int) ([]*entity.Product, error)
	Count(ctx context.Context) (int64, error)
	DecreaseStock(ctx context.Context, productID string, quantity int) error
	IncreaseStock(ctx context.Context, productID string, quantity int) error
	FindByStockRange(ctx context.Context, minStock, maxStock int) ([]*entity.Product, error)
}
