package usecase

import (
	"app-challenge/internal/domain/repository"
	"context"
	"errors"
	"time"
)

type ListOrdersRequest struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type ListOrderUserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type ListOrderItemResponse struct {
	ID        string          `json:"id"`
	Quantity  int             `json:"quantity"`
	Price     float64         `json:"price"`
	CreatedAt time.Time       `json:"createdAt"`
	Product   ProductResponse `json:"product"`
}

type ListOrderResponse struct {
	ID        string                  `json:"id"`
	Total     float64                 `json:"total"`
	CreatedAt time.Time               `json:"createdAt"`
	User      ListOrderUserResponse   `json:"user"`
	Items     []ListOrderItemResponse `json:"items"`
}

type ListOrdersResponse struct {
	Orders []ListOrderResponse `json:"orders"`
}

type ListOrdersUseCase struct {
	OrderRepo   repository.OrderRepository
	UserRepo    repository.UserRepository
	ProductRepo repository.ProductRepository
}

func (uc *ListOrdersUseCase) Execute(ctx context.Context, req ListOrdersRequest) (*ListOrdersResponse, error) {
	orders, err := uc.OrderRepo.List(ctx, req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}
	var resp []ListOrderResponse
	for _, o := range orders {
		user, err := uc.UserRepo.FindByID(ctx, o.UserID)
		if err != nil || user == nil {
			return nil, errors.New("user not found")
		}
		userResp := ListOrderUserResponse{
			ID:        user.ID.Value(),
			Name:      user.Name,
			Email:     user.Email.Value(),
			CreatedAt: user.CreatedAt,
		}
		var itemsResp []ListOrderItemResponse
		for _, item := range o.Items {
			var productResp ProductResponse
			product, _ := uc.ProductRepo.FindByID(ctx, item.ProductID)
			if product != nil {
				productResp = ProductResponse{
					ID:        product.ID.Value(),
					Name:      product.Name,
					Price:     product.Price.Amount(),
					Stock:     product.Stock,
					CreatedAt: product.CreatedAt,
				}
			}
			itemsResp = append(itemsResp, ListOrderItemResponse{
				ID:        item.ID.Value(),
				Quantity:  item.Quantity,
				Price:     item.Price.Amount(),
				CreatedAt: item.CreatedAt,
				Product:   productResp,
			})
		}
		orderResp := ListOrderResponse{
			ID:        o.ID.Value(),
			Total:     o.Total.Amount(),
			User:      userResp,
			Items:     itemsResp,
			CreatedAt: o.CreatedAt,
		}
		resp = append(resp, orderResp)
	}
	return &ListOrdersResponse{Orders: resp}, nil
}
