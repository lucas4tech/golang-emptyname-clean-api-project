package repository

import (
	"app-challenge/internal/domain/entity"
	"app-challenge/internal/domain/repository"
	"app-challenge/internal/domain/value_object"
	"app-challenge/internal/infra/persistence/gorm/model"
	"context"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) repository.ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) Create(ctx context.Context, product *entity.Product) error {
	m := &model.Product{
		ID:    product.ID.Value(),
		Name:  product.Name,
		Price: product.Price.Amount(),
		Stock: product.Stock,
	}
	return r.db.WithContext(ctx).Create(m).Error
}

func (r *ProductRepository) List(ctx context.Context, limit, offset int) ([]*entity.Product, error) {
	var products []model.Product
	err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&products).Error
	if err != nil {
		return nil, err
	}
	var result []*entity.Product
	for _, m := range products {
		idVO, _ := value_object.NewUUID(m.ID)
		priceVO, _ := value_object.NewMoney(int64(m.Price * 100))
		result = append(result, &entity.Product{
			ID:        idVO,
			Name:      m.Name,
			Stock:     m.Stock,
			Price:     priceVO,
			CreatedAt: m.CreatedAt,
		})
	}
	return result, nil
}

func (r *ProductRepository) FindByID(ctx context.Context, id string) (*entity.Product, error) {
	var m model.Product
	err := r.db.WithContext(ctx).First(&m, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	idVO, _ := value_object.NewUUID(m.ID)
	priceVO, _ := value_object.NewMoney(int64(m.Price * 100))
	return &entity.Product{
		ID:        idVO,
		Name:      m.Name,
		Stock:     m.Stock,
		Price:     priceVO,
		CreatedAt: m.CreatedAt,
	}, nil
}

func (r *ProductRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&model.Product{}).Count(&count).Error
	return count, err
}

func (r *ProductRepository) DecreaseStock(ctx context.Context, productID string, quantity int) error {
	return r.db.WithContext(ctx).
		Model(&model.Product{}).
		Where("id = ? AND stock >= ?", productID, quantity).
		Update("stock", gorm.Expr("stock - ?", quantity)).Error
}

func (r *ProductRepository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Delete(&model.Product{}, "id = ?", id).Error
}

func (r *ProductRepository) Update(ctx context.Context, product *entity.Product) error {
	return r.db.WithContext(ctx).
		Model(&model.Product{}).
		Where("id = ?", product.ID.Value()).
		Updates(map[string]interface{}{
			"name":  product.Name,
			"price": product.Price.Amount(),
			"stock": product.Stock,
		}).Error
}

func (r *ProductRepository) IncreaseStock(ctx context.Context, productID string, quantity int) error {
	return r.db.WithContext(ctx).
		Model(&model.Product{}).
		Where("id = ?", productID).
		Update("stock", gorm.Expr("stock + ?", quantity)).Error
}

func (r *ProductRepository) FindByStockRange(ctx context.Context, minStock, maxStock int) ([]*entity.Product, error) {
	var products []model.Product
	err := r.db.WithContext(ctx).
		Where("stock >= ? AND stock <= ?", minStock, maxStock).
		Find(&products).Error
	if err != nil {
		return nil, err
	}
	var result []*entity.Product
	for _, m := range products {
		idVO, _ := value_object.NewUUID(m.ID)
		priceVO, _ := value_object.NewMoney(int64(m.Price * 100))
		result = append(result, &entity.Product{
			ID:        idVO,
			Name:      m.Name,
			Stock:     m.Stock,
			Price:     priceVO,
			CreatedAt: m.CreatedAt,
		})
	}
	return result, nil
}
