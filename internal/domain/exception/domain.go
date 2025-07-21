package exception

import (
	"fmt"
)

type DomainError struct {
	Code    string
	Message string
	Details map[string]interface{}
}

func (e *DomainError) Error() string {
	if len(e.Details) > 0 {
		return fmt.Sprintf("[%s] %s", e.Code, e.Message)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}
