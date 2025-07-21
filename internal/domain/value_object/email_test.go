package value_object_test

import (
	"app-challenge/internal/domain/value_object"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEmail_Success(t *testing.T) {
	email, err := value_object.NewEmail("test@example.com")
	assert.NoError(t, err)
	assert.Equal(t, "test@example.com", email.Value())
}

func TestNewEmail_Invalid(t *testing.T) {
	email, err := value_object.NewEmail("invalid-email")
	assert.Error(t, err)
	assert.Nil(t, email)
}

func TestEmail_Equals(t *testing.T) {
	email1, _ := value_object.NewEmail("test@example.com")
	email2, _ := value_object.NewEmail("test@example.com")
	email3, _ := value_object.NewEmail("other@example.com")
	assert.True(t, email1.Equals(email2))
	assert.False(t, email1.Equals(email3))
}
