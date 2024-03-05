package validation

import "fmt"

type ServiceError struct {
	Message string `json:"message"`
}

func (e *ServiceError) Error() string {
	return e.Message
}

func NewUniqueValidationErrorMessage(value string) string {
	return fmt.Sprintf("'%s' is already taken", value)
}
