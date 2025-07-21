package value_object

import (
	"app-challenge/internal/domain/exception"

	"github.com/google/uuid"
)

type UUID struct {
	value string
}

func NewUUID(value string) (*UUID, error) {
	if value == "" {
		return nil, exception.NewValueObjectError("UUID", "value", "uuid cannot be empty")
	}
	_, err := uuid.Parse(value)
	if err != nil {
		return nil, exception.NewValueObjectError("UUID", "value", "invalid uuid format")
	}
	return &UUID{value: value}, nil
}

func NewUUIDv4() *UUID {
	return &UUID{value: uuid.NewString()}
}

func (u *UUID) Value() string {
	return u.value
}

func (u *UUID) String() string {
	return u.value
}

func (u *UUID) Equals(other *UUID) bool {
	if other == nil {
		return false
	}
	return u.value == other.value
}
