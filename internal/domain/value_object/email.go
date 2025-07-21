package value_object

import (
	"fmt"
	"regexp"
	"strings"

	"app-challenge/internal/domain/exception"
)

type Email struct {
	value string
}

func NewEmail(value string) (*Email, error) {
	if value == "" {
		return nil, exception.NewValueObjectError("Email", "value", "email cannot be empty")
	}

	normalizedEmail := strings.ToLower(strings.TrimSpace(value))
	if !isValidEmailFormat(normalizedEmail) {
		return nil, exception.NewValueObjectError("Email", "value", "valid email format")
	}

	return &Email{value: normalizedEmail}, nil
}

func (e *Email) Value() string {
	return e.value
}

func (e *Email) String() string {
	return e.value
}

func (e *Email) Equals(other *Email) bool {
	if other == nil {
		return false
	}
	return e.value == other.value
}

func isValidEmailFormat(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func (e *Email) Domain() string {
	parts := strings.Split(e.value, "@")
	if len(parts) != 2 {
		return ""
	}
	return parts[1]
}

func (e *Email) LocalPart() string {
	parts := strings.Split(e.value, "@")
	if len(parts) != 2 {
		return ""
	}
	return parts[0]
}

func (e *Email) IsValid() bool {
	return isValidEmailFormat(e.value)
}

func (e *Email) Format() string {
	return fmt.Sprintf("<%s>", e.value)
}
