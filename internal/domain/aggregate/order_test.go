package aggregate_test

import (
	"app-challenge/internal/domain/aggregate"
	"app-challenge/internal/domain/entity"
	"app-challenge/internal/domain/value_object"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOrder_Success(t *testing.T) {
	order, err := aggregate.NewOrder("user-1")
	assert.NoError(t, err)
	assert.Equal(t, "user-1", order.UserID)
	assert.NotNil(t, order.ID)
}

func TestNewOrder_EmptyUserID(t *testing.T) {
	order, err := aggregate.NewOrder("")
	assert.Error(t, err)
	assert.Nil(t, order)
}

func TestOrder_AddItem(t *testing.T) {
	order, _ := aggregate.NewOrder("user-1")
	price, _ := value_object.NewMoney(1000)
	item, _ := entity.NewOrderItem("prod-1", 2, price)
	err := order.AddItem(item)
	assert.NoError(t, err)
	assert.Len(t, order.Items, 1)
}

func TestOrder_AddItem_Duplicate(t *testing.T) {
	order, _ := aggregate.NewOrder("user-1")
	price, _ := value_object.NewMoney(1000)
	item1, _ := entity.NewOrderItem("prod-1", 2, price)
	item2, _ := entity.NewOrderItem("prod-1", 1, price)
	_ = order.AddItem(item1)
	err := order.AddItem(item2)
	assert.Error(t, err)
}

func TestOrder_CalculateTotal(t *testing.T) {
	order, _ := aggregate.NewOrder("user-1")
	price, _ := value_object.NewMoney(500)
	item1, _ := entity.NewOrderItem("prod-1", 2, price)
	item2, _ := entity.NewOrderItem("prod-2", 1, price)
	_ = order.AddItem(item1)
	_ = order.AddItem(item2)
	total, err := order.CalculateTotal()
	assert.NoError(t, err)
	assert.Equal(t, int64(1500), total.AmountInCents())
}
