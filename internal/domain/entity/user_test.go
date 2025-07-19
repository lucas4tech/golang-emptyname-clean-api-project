package entity_test

import (
	"app-challenge/internal/domain/entity"
	"app-challenge/internal/domain/value_object"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser_Success(t *testing.T) {
	email, _ := value_object.NewEmail("test@example.com")
	user, err := entity.NewUser("John Doe", email)
	assert.NoError(t, err)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, email, user.Email)
}

func TestNewUser_EmptyName(t *testing.T) {
	email, _ := value_object.NewEmail("test@example.com")
	user, err := entity.NewUser("", email)
	assert.Error(t, err)
	assert.Nil(t, user)
}

func TestNewUser_NilEmail(t *testing.T) {
	user, err := entity.NewUser("John Doe", nil)
	assert.Error(t, err)
	assert.Nil(t, user)
}

func TestUser_UpdateName(t *testing.T) {
	email, _ := value_object.NewEmail("test@example.com")
	user, _ := entity.NewUser("John Doe", email)
	err := user.UpdateName("Jane Doe")
	assert.NoError(t, err)
	assert.Equal(t, "Jane Doe", user.Name)
}

func TestUser_UpdateName_Invalid(t *testing.T) {
	email, _ := value_object.NewEmail("test@example.com")
	user, _ := entity.NewUser("John Doe", email)
	err := user.UpdateName("")
	assert.Error(t, err)
}

func TestUser_UpdateEmail(t *testing.T) {
	email, _ := value_object.NewEmail("test@example.com")
	user, _ := entity.NewUser("John Doe", email)
	newEmail, _ := value_object.NewEmail("new@example.com")
	err := user.UpdateEmail(newEmail)
	assert.NoError(t, err)
	assert.Equal(t, newEmail, user.Email)
}

func TestUser_UpdateEmail_Nil(t *testing.T) {
	email, _ := value_object.NewEmail("test@example.com")
	user, _ := entity.NewUser("John Doe", email)
	err := user.UpdateEmail(nil)
	assert.Error(t, err)
}
