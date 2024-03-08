package validation

import "fmt"

type ServiceError struct {
	Message string `json:"message"`
}

func (e *ServiceError) Error() string {
	return e.Message
}

type ResourceNotFoundError struct {
	Message string `json:"message"`
}

func (e *ResourceNotFoundError) Error() string {
	return e.Message
}

type NotModifiedError struct {
	Message string `json:"message"`
}

func (e *NotModifiedError) Error() string {
	return e.Message
}

func NewNotModifiedErrorMessage() string {
	return "Nothing to update."
}

func NewUniqueValidationErrorMessage(field string) string {
	return fmt.Sprintf("%s is already taken.", field)
}

func NewResourceNotFoundErrorMessage(resource string) string {
	return fmt.Sprintf("%s not found.", resource)
}

func NewInvalidCredentialsErrorMessage() string {
	return "invalid credentials."
}
