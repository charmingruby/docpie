package endpoints

import "fmt"

func NewCreateResponse(entity string) string {
	return fmt.Sprintf("%s created successfully.", entity)
}
