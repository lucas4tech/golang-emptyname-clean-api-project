package value_object_test

import (
	"app-challenge/internal/domain/value_object"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMoney_Success(t *testing.T) {
	m, err := value_object.NewMoney(1000)
	assert.NoError(t, err)
	assert.Equal(t, int64(1000), m.AmountInCents())
}

func TestNewMoney_Negative(t *testing.T) {
	m, err := value_object.NewMoney(-1)
	assert.Error(t, err)
	assert.Nil(t, m)
}

func TestMoney_Add(t *testing.T) {
	m1, _ := value_object.NewMoney(1000)
	m2, _ := value_object.NewMoney(500)
	res, err := m1.Add(m2)
	assert.NoError(t, err)
	assert.Equal(t, int64(1500), res.AmountInCents())
}

func TestMoney_Subtract(t *testing.T) {
	m1, _ := value_object.NewMoney(1000)
	m2, _ := value_object.NewMoney(500)
	res, err := m1.Subtract(m2)
	assert.NoError(t, err)
	assert.Equal(t, int64(500), res.AmountInCents())
}

func TestMoney_Subtract_Negative(t *testing.T) {
	m1, _ := value_object.NewMoney(500)
	m2, _ := value_object.NewMoney(1000)
	res, err := m1.Subtract(m2)
	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestMoney_Multiply(t *testing.T) {
	m, _ := value_object.NewMoney(1000)
	res, err := m.Multiply(2)
	assert.NoError(t, err)
	assert.Equal(t, int64(2000), res.AmountInCents())
}

func TestMoney_Equals(t *testing.T) {
	m1, _ := value_object.NewMoney(1000)
	m2, _ := value_object.NewMoney(1000)
	m3, _ := value_object.NewMoney(500)
	assert.True(t, m1.Equals(m2))
	assert.False(t, m1.Equals(m3))
}
