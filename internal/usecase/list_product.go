package usecase

import (
	"app-challenge/internal/domain/repository"
	"context"
	"time"
)

type ListProductsRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type ProductResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"createdAt"`
}

type ListProductsResponse struct {
	Products []ProductResponse `json:"products"`
}

type ListProductsUseCase struct {
	ProductRepo repository.ProductRepository
}

func (uc *ListProductsUseCase) Execute(ctx context.Context, req ListProductsRequest) (*ListProductsResponse, error) {
	products, err := uc.ProductRepo.List(ctx, req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}
	var resp []ProductResponse
	for _, p := range products {
		resp = append(resp, ProductResponse{
			ID:        p.ID.Value(),
			Name:      p.Name,
			Price:     p.Price.Amount(),
			Stock:     p.Stock,
			CreatedAt: p.CreatedAt,
		})
	}
	return &ListProductsResponse{Products: resp}, nil
}
