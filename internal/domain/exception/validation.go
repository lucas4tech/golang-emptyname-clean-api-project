package exception

import "fmt"

func NewRequiredFieldError(entity string, field string) *DomainError {
	return &DomainError{
		Code:    "REQUIRED_FIELD",
		Message: fmt.Sprintf("%s: field '%s' is required", entity, field),
		Details: map[string]interface{}{
			"entity": entity,
			"field":  field,
		},
	}
}

func NewInvalidFieldError(quantity int, reason string) *DomainError {
	return &DomainError{
		Code:    "INVALID_FIELD",
		Message: fmt.Sprintf("Invalid quantity %d: %s", quantity, reason),
		Details: map[string]interface{}{
			"quantity": quantity,
			"reason":   reason,
		},
	}
}

func NewEntityValidationError(entity string, field string, reason string) *DomainError {
	return &DomainError{
		Code:    "ENTITY_VALIDATION",
		Message: fmt.Sprintf("%s: %s - %s", entity, field, reason),
		Details: map[string]interface{}{
			"entity": entity,
			"field":  field,
			"reason": reason,
		},
	}
}

func NewDuplicateProductInOrderError() *DomainError {
	return &DomainError{
		Code:    "DUPLICATE_PRODUCT",
		Message: "duplicate product in order",
		Details: map[string]interface{}{
			"entity": "Order",
			"field":  "items",
		},
	}
}

func WrapWithContext(err error, context string) error {
	return fmt.Errorf("%s: %w", context, err)
}
