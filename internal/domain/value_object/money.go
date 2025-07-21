package value_object

import (
	"fmt"
	"math"

	"app-challenge/internal/domain/exception"
)

type Money struct {
	amount int64
}

func NewMoney(amountInCents int64) (*Money, error) {
	if amountInCents < 0 {
		return nil, exception.NewValueObjectError("Money", "amount", "amount cannot be negative")
	}

	return &Money{
		amount: amountInCents,
	}, nil
}

func (m *Money) AmountInCents() int64 {
	return m.amount
}

func (m *Money) Amount() float64 {
	return float64(m.amount) / 100.0
}

func (m *Money) String() string {
	return fmt.Sprintf("%.2f", m.Amount())
}

func (m *Money) Equals(other *Money) bool {
	if other == nil {
		return false
	}
	return m.amount == other.amount
}

func (m *Money) Add(other *Money) (*Money, error) {
	if other == nil {
		return nil, exception.NewValueObjectError("Money", "other", "cannot add nil money")
	}

	return NewMoney(m.amount + other.amount)
}

func (m *Money) Subtract(other *Money) (*Money, error) {
	if other == nil {
		return nil, exception.NewValueObjectError("Money", "other", "cannot subtract nil money")
	}

	if m.amount < other.amount {
		return nil, exception.NewValueObjectError("Money", "result", "result would be negative")
	}

	return NewMoney(m.amount - other.amount)
}

func (m *Money) Multiply(factor float64) (*Money, error) {
	if factor < 0 {
		return nil, exception.NewValueObjectError("Money", "factor", "factor cannot be negative")
	}

	newAmount := int64(math.Round(float64(m.amount) * factor))
	return NewMoney(newAmount)
}

func (m *Money) IsZero() bool {
	return m.amount == 0
}

func (m *Money) IsPositive() bool {
	return m.amount > 0
}

func (m *Money) IsNegative() bool {
	return m.amount < 0
}

func (m *Money) Format() string {
	return fmt.Sprintf("%.2f", m.Amount())
}
