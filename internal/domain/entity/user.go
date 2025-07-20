package entity

import (
	"strings"
	"time"

	"app-challenge/internal/domain/exception"
	"app-challenge/internal/domain/value_object"
)

type User struct {
	ID        *value_object.UUID
	Name      string
	Email     *value_object.Email
	CreatedAt time.Time
}

func NewUser(name string, email *value_object.Email) (*User, error) {
	if name == "" {
		return nil, exception.NewRequiredFieldError("User", "name")
	}

	if email == nil {
		return nil, exception.NewRequiredFieldError("User", "email")
	}

	if len(strings.TrimSpace(name)) < 2 {
		return nil, exception.NewEntityValidationError("User", "name", "name must have at least 2 characters")
	}

	now := time.Now()
	return &User{
		ID:        value_object.NewUUIDv4(),
		Name:      strings.TrimSpace(name),
		Email:     email,
		CreatedAt: now,
	}, nil
}

func (u *User) UpdateName(name string) error {
	if name == "" {
		return exception.NewRequiredFieldError("User", "name")
	}

	if len(strings.TrimSpace(name)) < 2 {
		return exception.NewEntityValidationError("User", "name", "name must have at least 2 characters")
	}

	u.Name = strings.TrimSpace(name)
	return nil
}

func (u *User) UpdateEmail(email *value_object.Email) error {
	if email == nil {
		return exception.NewRequiredFieldError("User", "email")
	}
	u.Email = email
	return nil
}

func (u *User) GetEmail() string {
	return u.Email.Value()
}

func (u *User) GetName() string {
	return u.Name
}
