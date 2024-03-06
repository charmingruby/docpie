package validation

import (
	"fmt"

	"github.com/charmingruby/upl/helpers"
)

type EndpointError struct {
	Message string `json:"message"`
}

func (e *EndpointError) Error() string {
	return e.Message
}

func NewPayloadErrorMessage(requiredFields []string) string {
	var fieldsStr string

	for idx, field := range requiredFields {
		if idx == 0 {
			fieldsStr += field
		} else {
			if idx+1 == len(requiredFields) {
				fieldsStr += fmt.Sprintf(" and %s", field)
			} else {
				fieldsStr += fmt.Sprintf(", %s", field)
			}
		}
	}

	statementConnector := helpers.If[string](len(requiredFields) <= 1, "is", "are")

	return fmt.Sprintf("Invalid payload, %s %s required.", fieldsStr, statementConnector)
}

func NewEmptyPayloadFieldsErrorMessage(fields []string) string {
	var fieldsStr string

	if len(fields) == 1 {
		return fmt.Sprintf("Invalid payload, %s cannot be blank.", fields[0])
	}

	for idx, field := range fields {
		if idx+1 == len(fields) {
			fieldsStr += fmt.Sprintf(" and %s", field)
		} else {
			if idx == 0 {
				fieldsStr += field
			} else {
				fieldsStr += fmt.Sprintf(", %s", field)
			}
		}
	}

	return fmt.Sprintf("Invalid payload, %s cannot be blank.", fieldsStr)

}
