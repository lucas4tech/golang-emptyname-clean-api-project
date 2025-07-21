package usecase

import (
	"app-challenge/internal/domain/repository"
	"context"
	"time"
)

type ListUsersWithOrdersUseCase struct {
	UserRepo    repository.UserRepository
	ProductRepo repository.ProductRepository
}

type ListUsersWithOrdersRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type ListUsersWithOrdersResponse struct {
	Users []UserWithOrdersResponse `json:"users"`
}

type UserWithOrdersResponse struct {
	ID        string          `json:"id"`
	Name      string          `json:"name"`
	Email     string          `json:"email"`
	Orders    []OrderResponse `json:"orders"`
	CreatedAt time.Time       `json:"createdAt"`
}

type OrderResponse struct {
	ID        string              `json:"id"`
	Total     float64             `json:"total"`
	Items     []OrderItemResponse `json:"items"`
	CreatedAt time.Time           `json:"createdAt"`
}

type OrderProductResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"createdAt"`
}

type OrderItemResponse struct {
	ID        string               `json:"id"`
	Quantity  int                  `json:"quantity"`
	Product   OrderProductResponse `json:"product"`
	Price     float64              `json:"price"`
	CreatedAt time.Time            `json:"createdAt"`
}

func (uc *ListUsersWithOrdersUseCase) Execute(ctx context.Context, req ListUsersWithOrdersRequest) (*ListUsersWithOrdersResponse, error) {
	users, err := uc.UserRepo.List(ctx, req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}
	var resp []UserWithOrdersResponse
	for _, u := range users {
		userResp := UserWithOrdersResponse{
			ID:        u.ID.Value(),
			Name:      u.Name,
			Email:     u.Email.Value(),
			CreatedAt: u.CreatedAt,
		}
		resp = append(resp, userResp)
	}
	return &ListUsersWithOrdersResponse{Users: resp}, nil
}
