package entity_test

import (
	"app-challenge/internal/domain/entity"
	"app-challenge/internal/domain/value_object"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct_Success(t *testing.T) {
	price, _ := value_object.NewMoney(1000)
	product, err := entity.NewProduct("Notebook", price, 10)
	assert.NoError(t, err)
	assert.Equal(t, "Notebook", product.Name)
	assert.Equal(t, 10, product.Stock)
	assert.True(t, product.Price.Equals(price))
}

func TestNewProduct_EmptyName(t *testing.T) {
	price, _ := value_object.NewMoney(1000)
	product, err := entity.NewProduct("", price, 10)
	assert.Error(t, err)
	assert.Nil(t, product)
}

func TestNewProduct_NilPrice(t *testing.T) {
	product, err := entity.NewProduct("Notebook", nil, 10)
	assert.Error(t, err)
	assert.Nil(t, product)
}

func TestNewProduct_NegativeStock(t *testing.T) {
	price, _ := value_object.NewMoney(1000)
	product, err := entity.NewProduct("Notebook", price, -1)
	assert.Error(t, err)
	assert.Nil(t, product)
}

func TestProduct_UpdateName(t *testing.T) {
	price, _ := value_object.NewMoney(1000)
	product, _ := entity.NewProduct("Notebook", price, 10)
	err := product.UpdateName("Mouse")
	assert.NoError(t, err)
	assert.Equal(t, "Mouse", product.Name)
}

func TestProduct_UpdatePrice(t *testing.T) {
	price, _ := value_object.NewMoney(1000)
	product, _ := entity.NewProduct("Notebook", price, 10)
	newPrice, _ := value_object.NewMoney(2000)
	err := product.UpdatePrice(newPrice)
	assert.NoError(t, err)
	assert.True(t, product.Price.Equals(newPrice))
}

func TestProduct_IncreaseDecreaseStock(t *testing.T) {
	price, _ := value_object.NewMoney(1000)
	product, _ := entity.NewProduct("Notebook", price, 10)
	err := product.IncreaseStock(5)
	assert.NoError(t, err)
	assert.Equal(t, 15, product.Stock)
	err = product.DecreaseStock(3)
	assert.NoError(t, err)
	assert.Equal(t, 12, product.Stock)
}

func TestProduct_DecreaseStock_Insufficient(t *testing.T) {
	price, _ := value_object.NewMoney(1000)
	product, _ := entity.NewProduct("Notebook", price, 2)
	err := product.DecreaseStock(5)
	assert.Error(t, err)
}
