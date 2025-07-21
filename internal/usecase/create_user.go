package usecase

import (
	"app-challenge/internal/domain/entity"
	"app-challenge/internal/domain/repository"
	"app-challenge/internal/domain/value_object"
	"app-challenge/pkg/uow"
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type CreateUserRequest struct {
	Name  string
	Email string
}

type CreateUserResponse struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateUserUseCase struct {
	UserRepo repository.UserRepository
	Uow      *uow.UnitOfWork
}

func (uc *CreateUserUseCase) Execute(ctx context.Context, req CreateUserRequest) (resp *CreateUserResponse, err error) {
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
	emailVO, err := value_object.NewEmail(req.Email)
	if err != nil {
		return nil, err
	}
	user, err := entity.NewUser(req.Name, emailVO)
	if err != nil {
		return nil, err
	}

	existing, err := uc.UserRepo.FindByEmail(ctx, req.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if existing != nil {
		return nil, errors.New("email already registered")
	}

	err = uc.UserRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return &CreateUserResponse{
		ID:        user.ID.Value(),
		Name:      user.Name,
		Email:     user.Email.Value(),
		CreatedAt: user.CreatedAt,
	}, nil
}
