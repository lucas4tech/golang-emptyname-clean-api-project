package repository

import (
	"app-challenge/internal/domain/aggregate"
	"app-challenge/internal/domain/entity"
	"context"
)

type OrderRepository interface {
	Create(ctx context.Context, order *aggregate.Order) error
	FindByID(ctx context.Context, id string) (*aggregate.Order, error)
	Update(ctx context.Context, order *aggregate.Order) error
	List(ctx context.Context, limit, offset int) ([]*aggregate.Order, error)
	Delete(ctx context.Context, id string) error
	FindByUserID(ctx context.Context, userID string, limit, offset int) ([]*aggregate.Order, error)
	CountByUserID(ctx context.Context, userID string) (int64, error)
	WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error
	CreateOrderItem(ctx context.Context, item *entity.OrderItem) error
}
