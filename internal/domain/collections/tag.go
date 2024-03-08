package collections

import (
	"time"

	"github.com/charmingruby/upl/internal/core"
	"github.com/charmingruby/upl/internal/validation"
	"github.com/charmingruby/upl/internal/validation/errs"
)

func NewCollectionTag(name, description string) (*CollectionTag, error) {
	tag := CollectionTag{
		ID:          core.NewId(),
		Name:        name,
		Description: description,
		CreatedAt:   time.Now(),
		UpdatedAt:   nil,
	}

	if err := tag.Validate(); err != nil {
		return nil, err
	}

	return &tag, nil
}

type CollectionTag struct {
	ID          string     `db:"id" json:"id"`
	Name        string     `db:"name" json:"name"`
	Description string     `db:"description" json:"description"`
	CreatedAt   time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updated_at"`
}

func (c *CollectionTag) Validate() error {
	if validation.IsEmpty(c.Name) {
		return &errs.ValidationError{
			Message: errs.EntitieisRequiredFieldErrorMessage("name"),
		}
	}

	if validation.IsGreater(c.Name, 20) {
		return &errs.ValidationError{
			Message: errs.EntitiesFieldLengthErrorMessage("name", 20, false),
		}
	}

	if validation.IsLower(c.Name, 2) {
		return &errs.ValidationError{
			Message: errs.EntitiesFieldLengthErrorMessage("name", 2, true),
		}
	}

	if validation.IsEmpty(c.Description) {
		return &errs.ValidationError{
			Message: errs.EntitieisRequiredFieldErrorMessage("description"),
		}
	}

	if validation.IsGreater(c.Description, 32) {
		return &errs.ValidationError{
			Message: errs.EntitiesFieldLengthErrorMessage("description", 32, false),
		}
	}

	if validation.IsLower(c.Description, 8) {
		return &errs.ValidationError{
			Message: errs.EntitiesFieldLengthErrorMessage("description", 8, true),
		}
	}

	return nil
}
