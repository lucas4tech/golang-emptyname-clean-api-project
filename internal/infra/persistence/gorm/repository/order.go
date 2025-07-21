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

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) repository.OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(ctx context.Context, order *aggregate.Order) error {
	m := &model.Order{
		ID:     order.ID.Value(),
		UserID: order.UserID,
		Total:  order.Total.Amount(),
	}
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *OrderRepository) List(ctx context.Context, limit, offset int) ([]*aggregate.Order, error) {
	var orders []model.Order
	err := r.db.WithContext(ctx).
		Limit(limit).Offset(offset).
		Find(&orders).Error
	if err != nil {
		return nil, err
	}
	var result []*aggregate.Order
	for _, m := range orders {
		idVO, _ := value_object.NewUUID(m.ID)
		totalVO, _ := value_object.NewMoney(int64(m.Total * 100))
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
			Total:     totalVO,
			Items:     items,
			CreatedAt: m.CreatedAt,
		})
	}
	return result, nil
}

func (r *OrderRepository) FindByID(ctx context.Context, id string) (*aggregate.Order, error) {
	var m model.Order
	err := r.db.WithContext(ctx).First(&m, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	idVO, _ := value_object.NewUUID(m.ID)
	totalVO, _ := value_object.NewMoney(int64(m.Total * 100))
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
	return &aggregate.Order{
		ID:        idVO,
		UserID:    m.UserID,
		Total:     totalVO,
		Items:     items,
		CreatedAt: m.CreatedAt,
	}, nil
}

func (r *OrderRepository) CountByUserID(ctx context.Context, userID string) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.Order{}).Where("user_id = ?", userID).Count(&count).Error
	return count, err
}

func (r *OrderRepository) Update(ctx context.Context, order *aggregate.Order) error {
	return r.db.WithContext(ctx).
		Model(&model.Order{}).
		Where("id = ?", order.ID.Value()).
		Updates(map[string]interface{}{
			"user_id": order.UserID,
			"total":   order.Total.Amount(),
		}).Error
}

func (r *OrderRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&model.Order{}, "id = ?", id).Error
}

func (r *OrderRepository) FindByUserID(ctx context.Context, userID string, limit, offset int) ([]*aggregate.Order, error) {
	var orders []model.Order
	err := r.db.WithContext(ctx).
		Where("user_id = ?", userID).
		Limit(limit).Offset(offset).
		Find(&orders).Error
	if err != nil {
		return nil, err
	}
	var result []*aggregate.Order
	for _, m := range orders {
		idVO, _ := value_object.NewUUID(m.ID)
		totalVO, _ := value_object.NewMoney(int64(m.Total * 100))
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
			Total:     totalVO,
			Items:     items,
			CreatedAt: m.CreatedAt,
		})
	}
	return result, nil
}

func (r *OrderRepository) CreateOrderItem(ctx context.Context, item *entity.OrderItem) error {
	m := &model.OrderItem{
		ID:        item.ID.Value(),
		OrderID:   item.OrderID,
		ProductID: item.ProductID,
		Quantity:  item.Quantity,
		Price:     item.Price.Amount(),
	}
	return r.db.WithContext(ctx).Create(m).Error
}

type txKeyType struct{}

var txKey = txKeyType{}

func (r *OrderRepository) WithTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		newCtx := context.WithValue(ctx, txKey, tx)
		return fn(newCtx)
	})
}
