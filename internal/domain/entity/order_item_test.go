package entity_test

import (
	"testing"

	"app-challenge/internal/domain/entity"
	"app-challenge/internal/domain/value_object"

	"github.com/stretchr/testify/assert"
)

func TestNewOrderItem_Success(t *testing.T) {
	price, _ := value_object.NewMoney(1000)
	item, err := entity.NewOrderItem("prod-1", 2, price)
	assert.NoError(t, err)
	assert.Equal(t, "prod-1", item.ProductID)
	assert.Equal(t, 2, item.Quantity)
	assert.True(t, item.Price.Equals(price))
}

func TestNewOrderItem_InvalidQuantity(t *testing.T) {
	price, _ := value_object.NewMoney(1000)
	item, err := entity.NewOrderItem("prod-1", 0, price)
	assert.Error(t, err)
	assert.Nil(t, item)
}

func TestNewOrderItem_NilPrice(t *testing.T) {
	item, err := entity.NewOrderItem("prod-1", 1, nil)
	assert.Error(t, err)
	assert.Nil(t, item)
}

func TestOrderItem_GetTotalPrice(t *testing.T) {
	price, _ := value_object.NewMoney(500)
	item, _ := entity.NewOrderItem("prod-1", 3, price)
	total, err := item.GetTotalPrice()
	assert.NoError(t, err)
	assert.Equal(t, int64(1500), total.AmountInCents())
}
