package endpoints

import "fmt"

func CreatedResponse(entity string) string {
	return fmt.Sprintf("%s created successfully.", entity)
}

func UpdatedResponse(entity string) string {
	return fmt.Sprintf("%s updated successfully.", entity)
}

func ModifiedResponse(entity, field string) string {
	return fmt.Sprintf("%s %s modified successfully.", entity, field)
}

func DeleteResponse(entity string) string {
	return fmt.Sprintf("%s deleted successfully.", entity)

}
