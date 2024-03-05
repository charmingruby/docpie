package validation

import "fmt"

type HTTPError struct {
	Message string `json:"message"`
}

func (e *HTTPError) Error() string {
	return e.Message
}

func NewPayloadErrorResponse(requiredFields []string) string {
	var fieldsStr string

	for idx, field := range requiredFields {
		if idx+1 == len(requiredFields) {
			fieldsStr += fmt.Sprintf(" and %s", field)
		} else {
			if idx == 0 {
				fieldsStr += field
			} else {
				fieldsStr += fmt.Sprintf(", %s", field)
			}
		}
	}

	return fmt.Sprintf("Invalid payload, please provide at least: %s.", fieldsStr)
}
