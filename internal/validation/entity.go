package validation

import (
	"fmt"

	"github.com/charmingruby/upl/helpers"
)

type ValidationError struct {
	Message string `json:"message"`
}

func (e *ValidationError) Error() string {
	return e.Message
}

func NewRequiredFieldErrorMessage(field string) string {
	return fmt.Sprintf("%s cannot be blank", field)
}

func NewFieldLengthErrorMessage(field string, quantity int, isMinError bool) string {
	minMsg := fmt.Sprintf("%s must be at least %d characters", field, quantity)
	maxMsg := fmt.Sprintf("%s must be a maximum of %d characters", field, quantity)

	msg := helpers.If[string](
		isMinError,
		minMsg,
		maxMsg,
	)

	return msg
}
