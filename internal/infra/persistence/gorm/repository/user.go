package repository

import (
	"app-challenge/internal/domain/aggregate"
	"app-challenge/internal/domain/entity"
	"app-challenge/internal/domain/repository"
	"app-challenge/internal/domain/value_object"
	"app-challenge/internal/infra/persistence/gorm/model"
	"context"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *entity.User) error {
	m := &model.User{
		ID:    user.ID.Value(),
		Name:  user.Name,
		Email: user.Email.Value(),
	}
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *UserRepository) List(ctx context.Context, limit, offset int) ([]*entity.User, error) {
	var users []model.User
	err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&users).Error
	if err != nil {
		return nil, err
	}
	var result []*entity.User
	for _, m := range users {
		idVO, err := value_object.NewUUID(m.ID)
		if err != nil {
			return nil, err
		}
		emailVO, _ := value_object.NewEmail(m.Email)
		result = append(result, &entity.User{
			ID:        idVO,
			Name:      m.Name,
			Email:     emailVO,
			CreatedAt: m.CreatedAt,
		})
	}
	return result, nil
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*entity.User, error) {
	var m model.User
	err := r.db.WithContext(ctx).First(&m, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	idVO, _ := value_object.NewUUID(m.ID)
	emailVO, _ := value_object.NewEmail(m.Email)
	return &entity.User{
		ID:        idVO,
		Name:      m.Name,
		Email:     emailVO,
		CreatedAt: m.CreatedAt,
	}, nil
}

func (r *UserRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.User{}).Count(&count).Error
	return count, err
}

func (r *UserRepository) ListOrdersByUserID(ctx context.Context, userID string) ([]*aggregate.Order, error) {
	var orders []model.Order
	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Find(&orders).Error
	if err != nil {
		return nil, err
	}
	var result []*aggregate.Order
	for _, m := range orders {
		idVO, _ := value_object.NewUUID(m.ID)
		var itemModels []model.OrderItem
		err = r.db.WithContext(ctx).Where("order_id = ?", m.ID).Find(&itemModels).Error
		if err != nil {
			return nil, err
		}
		var items []*entity.OrderItem
		for _, itemModel := range itemModels {
			itemID, _ := value_object.NewUUID(itemModel.ID)
			priceVO, _ := value_object.NewMoney(int64(itemModel.Price * 100))
			items = append(items, &entity.OrderItem{
				ID:        itemID,
				ProductID: itemModel.ProductID,
				OrderID:   itemModel.OrderID,
				Quantity:  itemModel.Quantity,
				Price:     priceVO,
				CreatedAt: itemModel.CreatedAt,
			})
		}
		result = append(result, &aggregate.Order{
			ID:        idVO,
			UserID:    m.UserID,
			Items:     items,
			CreatedAt: m.CreatedAt,
		})
	}
	return result, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var m model.User
	err := r.db.WithContext(ctx).First(&m, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	idVO, err := value_object.NewUUID(m.ID)
	if err != nil {
		return nil, err
	}
	emailVO, err := value_object.NewEmail(m.Email)
	if err != nil {
		return nil, err
	}
	return &entity.User{
		ID:        idVO,
		Name:      m.Name,
		Email:     emailVO,
		CreatedAt: m.CreatedAt,
	}, nil
}
