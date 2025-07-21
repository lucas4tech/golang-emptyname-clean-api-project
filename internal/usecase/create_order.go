package usecase

import (
	"app-challenge/internal/domain/aggregate"
	"app-challenge/internal/domain/entity"
	"app-challenge/internal/domain/repository"
	"app-challenge/internal/domain/value_object"
	"app-challenge/pkg/uow"
	"context"
	"errors"
	"time"
)

type CreateOrderRequest struct {
	UserID string
	Items  []CreateOrderItemRequest
}

type CreateOrderItemRequest struct {
	ProductID string
	Quantity  int
}

type CreateOrderItemProductResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateOrderItemResponse struct {
	ID        string                         `json:"id"`
	Product   CreateOrderItemProductResponse `json:"product"`
	Quantity  int                            `json:"quantity"`
	Price     float64                        `json:"price"`
	CreatedAt time.Time                      `json:"createdAt"`
}

type CreateOrderUserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateOrderResponse struct {
	ID        string                    `json:"id"`
	Total     float64                   `json:"total"`
	CreatedAt time.Time                 `json:"createdAt"`
	User      CreateOrderUserResponse   `json:"user"`
	Items     []CreateOrderItemResponse `json:"items"`
}

type CreateOrderUseCase struct {
	OrderRepo   repository.OrderRepository
	ProductRepo repository.ProductRepository
	UserRepo    repository.UserRepository
	Uow         *uow.UnitOfWork
}

func (uc *CreateOrderUseCase) Execute(ctx context.Context, req CreateOrderRequest) (resp *CreateOrderResponse, err error) {
	err = uc.Uow.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			uc.Uow.Rollback()
		} else {
			uc.Uow.Commit()
		}
	}()
	user, err := uc.UserRepo.FindByID(ctx, req.UserID)
	if err != nil || user == nil {
		return nil, errors.New("user not found")
	}

	if len(req.Items) == 0 {
		return nil, errors.New("order must have at least one item")
	}

	var orderItems []*entity.OrderItem
	var total int64

	for _, item := range req.Items {
		product, err := uc.ProductRepo.FindByID(ctx, item.ProductID)
		if err != nil || product == nil {
			return nil, errors.New("product not found: " + item.ProductID)
		}

		if product.Stock < item.Quantity {
			return nil, errors.New("insufficient stock for product: " + product.Name)
		}

		product.Stock -= item.Quantity
		orderItem, err := entity.NewOrderItem(item.ProductID, item.Quantity, product.Price)
		if err != nil {
			return nil, err
		}

		if err := uc.ProductRepo.Update(ctx, product); err != nil {
			return nil, err
		}

		orderItems = append(orderItems, orderItem)
		total += int64(product.Price.Amount() * float64(item.Quantity) * 100)
	}

	order, err := aggregate.NewOrder(req.UserID)
	if err != nil {
		return nil, err
	}

	for _, item := range orderItems {
		if err := order.AddItem(item); err != nil {
			return nil, err
		}
	}

	totalMoney, err := value_object.NewMoney(total)
	if err != nil {
		return nil, err
	}
	order.Total = totalMoney

	if err := uc.OrderRepo.Create(ctx, order); err != nil {
		return nil, err
	}

	for _, item := range order.Items {
		item.OrderID = order.ID.Value()
		if err := uc.OrderRepo.CreateOrderItem(ctx, item); err != nil {
			return nil, err
		}
	}

	var itemsResp []CreateOrderItemResponse
	for _, item := range order.Items {
		product, _ := uc.ProductRepo.FindByID(ctx, item.ProductID)
		var productResp CreateOrderItemProductResponse
		if product != nil {
			productResp = CreateOrderItemProductResponse{
				ID:        product.ID.Value(),
				Name:      product.Name,
				Price:     product.Price.Amount(),
				Stock:     product.Stock,
				CreatedAt: product.CreatedAt,
			}
		}
		itemsResp = append(itemsResp, CreateOrderItemResponse{
			ID:        item.ID.Value(),
			Product:   productResp,
			Quantity:  item.Quantity,
			Price:     item.Price.Amount(),
			CreatedAt: item.CreatedAt,
		})
	}

	userResp := CreateOrderUserResponse{
		ID:        user.ID.Value(),
		Name:      user.Name,
		Email:     user.Email.Value(),
		CreatedAt: user.CreatedAt,
	}

	return &CreateOrderResponse{
		ID:        order.ID.Value(),
		Total:     order.Total.Amount(),
		CreatedAt: order.CreatedAt,
		User:      userResp,
		Items:     itemsResp,
	}, nil
}
