package endpoints

import "fmt"

func NewCreateResponse(identifier string) string {
	return fmt.Sprintf("'%s' created successfully.", identifier)
}
