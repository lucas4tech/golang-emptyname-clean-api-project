package usecase

import (
	"app-challenge/internal/domain/entity"
	"app-challenge/internal/domain/repository"
	"app-challenge/internal/domain/value_object"
	"app-challenge/pkg/uow"
	"context"
	"time"
)

type CreateProductRequest struct {
	Name  string
	Price float64
	Stock int
}

type CreateProductResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateProductUseCase struct {
	ProductRepo repository.ProductRepository
	Uow         *uow.UnitOfWork
}

func (uc *CreateProductUseCase) Execute(ctx context.Context, req CreateProductRequest) (resp *CreateProductResponse, err error) {
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
	money, err := value_object.NewMoney(int64(req.Price * 100))
	if err != nil {
		return nil, err
	}
	product, err := entity.NewProduct(req.Name, money, req.Stock)
	if err != nil {
		return nil, err
	}
	err = uc.ProductRepo.Create(ctx, product)
	if err != nil {
		return nil, err
	}
	return &CreateProductResponse{
		ID:        product.ID.Value(),
		Name:      product.Name,
		Price:     product.Price.Amount(),
		Stock:     product.Stock,
		CreatedAt: product.CreatedAt,
	}, nil
}
