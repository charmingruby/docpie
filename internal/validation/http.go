package validation

import "fmt"

type EndpointError struct {
	Message string `json:"message"`
}

func (e *EndpointError) Error() string {
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

	return fmt.Sprintf("Invalid payload, %s are required.", fieldsStr)
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
