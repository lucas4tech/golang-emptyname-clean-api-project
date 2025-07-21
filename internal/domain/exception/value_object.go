package exception

import "fmt"

func NewValueObjectError(valueObject string, field string, reason string) *DomainError {
	return &DomainError{
		Code:    "VALUE_OBJECT_ERROR",
		Message: fmt.Sprintf("%s value object validation failed: %s - %s", valueObject, field, reason),
		Details: map[string]interface{}{
			"valueobject": valueObject,
			"field":       field,
			"reason":      reason,
		},
	}
}
