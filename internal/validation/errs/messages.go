package errs

import (
	"fmt"

	"github.com/charmingruby/upl/helpers"
)

///////////////////
// Entity        //
///////////////////
func EntitieisRequiredFieldErrorMessage(field string) string {
	return fmt.Sprintf("%s cannot be blank", field)
}

func EntitiesFieldLengthErrorMessage(field string, quantity int, isMinError bool) string {
	minMsg := fmt.Sprintf("%s must be at least %d characters", field, quantity)
	maxMsg := fmt.Sprintf("%s must be a maximum of %d characters", field, quantity)

	msg := helpers.If[string](
		isMinError,
		minMsg,
		maxMsg,
	)

	return msg
}

///////////////////
// File          //
///////////////////
func FilesInvalidMimetypeErrorMessage(mimetypeUnmatched string, validMimetypes []string) string {
	var msg = fmt.Sprintf(".%s is not a valid mimetype, please provide a valid mimetype: ", mimetypeUnmatched)

	for idx, mimetype := range validMimetypes {
		if idx+1 == len(mimetypeUnmatched) {
			msg += fmt.Sprintf("or .%s.", mimetype)
			continue
		}

		msg += fmt.Sprintf(".%s, ", mimetype)
	}

	return msg
}

func FilesReachesMaximumSizeErrorMessage(currentSize, maxSize int64) string {
	return fmt.Sprintf("%d perpasses the limit of %d bytes", currentSize, maxSize)
}

func FilesNoFileErrorMessage(key string) string {
	return fmt.Sprintf("no file found for key '%s'", key)
}

///////////////////
// HTTP          //
///////////////////
func HTTPPayloadErrorMessage(requiredFields []string) string {
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

func HTTPEmptyPayloadFieldsErrorMessage(fields []string) string {
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

///////////////////
// Database       //
///////////////////
func DatabaseRepositoryNotAbleErrorMessage(repositoryName string) string {
	return fmt.Sprintf("%s repository wasn't able to build all the statements", repositoryName)
}

func DatabaseQueryPreparationErrorMessage(queryName, err string) string {
	return fmt.Sprintf("Error preparing statement %s: %s", queryName, err)
}

func DatabaseQueryNotPreparedErrorMessage(query string) string {
	return fmt.Sprintf("prepared statement '%s' not found", query)
}

// From
func DatabaseQueryErrorMessage(entity, action string, err error) string {
	return fmt.Sprintf("error %s %s: %v", action, entity, err)
}

func DatabaseResourceNotFoundErrorMessage(entity string) string {
	return fmt.Sprintf("%s not found.", entity)
}

// To

///////////////////
// Service       //
///////////////////
func ServicesNotModifiedErrorMessage() string {
	return "Nothing to update."
}

func ServicesUniqueValidationErrorMessage(field string) string {
	return fmt.Sprintf("%s is already taken.", field)
}

func ServicesResourceNotFoundErrorMessage(resource string) string {
	return fmt.Sprintf("%s not found.", resource)
}

func ServicesInvalidCredentialsErrorMessage() string {
	return "Invalid credentials."
}
