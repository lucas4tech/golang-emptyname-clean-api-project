package aggregate

import (
	"time"

	"app-challenge/internal/domain/entity"
	"app-challenge/internal/domain/exception"
	"app-challenge/internal/domain/value_object"
)

type Order struct {
	ID        *value_object.UUID
	UserID    string
	Items     []*entity.OrderItem
	Total     *value_object.Money
	CreatedAt time.Time
}

func NewOrder(userID string) (*Order, error) {
	if userID == "" {
		return nil, exception.NewRequiredFieldError("Order", "userID")
	}

	now := time.Now()
	return &Order{
		ID:        value_object.NewUUIDv4(),
		UserID:    userID,
		Items:     []*entity.OrderItem{},
		CreatedAt: now,
	}, nil
}

func (o *Order) AddItem(item *entity.OrderItem) error {
	if item == nil {
		return exception.NewRequiredFieldError("Order", "item")
	}

	for _, existingItem := range o.Items {
		if existingItem.ProductID == item.ProductID {
			return exception.NewDuplicateProductInOrderError()
		}
	}

	o.Items = append(o.Items, item)
	return nil
}

func (o *Order) CalculateTotal() (*value_object.Money, error) {
	if len(o.Items) == 0 {
		return value_object.NewMoney(0)
	}

	var total *value_object.Money
	for i, item := range o.Items {
		itemTotal, err := item.GetTotalPrice()
		if err != nil {
			return nil, exception.WrapWithContext(err, "error calculating item total")
		}
		if i == 0 {
			total = itemTotal
		} else {
			total, err = total.Add(itemTotal)
			if err != nil {
				return nil, exception.WrapWithContext(err, "error adding item total")
			}
		}
	}
	return total, nil
}
