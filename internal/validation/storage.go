package validation

import "fmt"

type StorageError struct {
	Message string `json:"error"`
}

func (e *StorageError) Error() string {
	return e.Message
}

func NewRepositoryStatementsPreparationErrorMessage(repositoryName string) string {
	return fmt.Sprintf("%s wasn't able to build all the statements", repositoryName)
}

func NewQueryStatementPreparationErrorMessage(query string) string {
	return fmt.Sprintf("prepared statement '%s' not found", query)
}

func NewQueryErrorMessage(entity, action string, err error) string {
	return fmt.Sprintf("error %s %s: %v", action, entity, err)
}

func NewResourceNotFoundByErrorMessage(identifier string, entity, field string) string {
	return fmt.Sprintf("error finding %s with %s: '%s'", entity, field, identifier)
}
