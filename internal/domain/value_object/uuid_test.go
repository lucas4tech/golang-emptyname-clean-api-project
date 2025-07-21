package value_object_test

import (
	"app-challenge/internal/domain/value_object"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewUUIDv4(t *testing.T) {
	u := value_object.NewUUIDv4()
	assert.NotNil(t, u)
	assert.NotEmpty(t, u.Value())
}

func TestNewUUID_Success(t *testing.T) {
	id := uuid.New().String()
	u, err := value_object.NewUUID(id)
	assert.NoError(t, err)
	assert.Equal(t, id, u.Value())
}

func TestNewUUID_Invalid(t *testing.T) {
	u, err := value_object.NewUUID("invalid-uuid")
	assert.Error(t, err)
	assert.Nil(t, u)
}

func TestUUID_Equals(t *testing.T) {
	id := uuid.New().String()
	u1, _ := value_object.NewUUID(id)
	u2, _ := value_object.NewUUID(id)
	u3 := value_object.NewUUIDv4()
	assert.True(t, u1.Equals(u2))
	assert.False(t, u1.Equals(u3))
}
