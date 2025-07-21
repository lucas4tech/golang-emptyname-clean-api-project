package repository

import (
	"app-challenge/internal/domain/aggregate"
	"app-challenge/internal/domain/entity"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) error
	List(ctx context.Context, limit, offset int) ([]*entity.User, error)
	FindByID(ctx context.Context, id string) (*entity.User, error)
	ListOrdersByUserID(ctx context.Context, userID string) ([]*aggregate.Order, error)
	FindByEmail(ctx context.Context, email string) (*entity.User, error)
}
