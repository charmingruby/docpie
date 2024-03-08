package collections

import (
	"time"

	"github.com/charmingruby/upl/helpers"
	"github.com/charmingruby/upl/internal/core"
	"github.com/charmingruby/upl/internal/validation"
	"github.com/charmingruby/upl/pkg/cryptography"
)

func NewCollection(name, secret, description, tag, creatorID string) (*Collection, error) {
	collection := &Collection{
		ID:              core.NewId(),
		Name:            name,
		Description:     helpers.IfOrNil[string](description != "", description),
		Secret:          secret,
		TagID:           nil,
		Tag:             tag,
		CreatorID:       creatorID,
		DeletedBy:       nil,
		UploadsQuantity: 0,
		MembersQuantity: 0,
		CreatedAt:       time.Now(),
		UpdatedAt:       nil,
		DeletedAt:       nil,
	}

	if err := collection.encryptSecret(); err != nil {
		return nil, err
	}

	return collection, nil
}

type Collection struct {
	ID              string     `db:"id" json:"id"`
	Name            string     `db:"name" json:"name"`
	Description     *string    `db:"description" json:"description"`
	Secret          string     `db:"secret" json:"secret"`
	Tag             string     `db:"tag" json:"tag"`
	UploadsQuantity uint       `db:"uploads_quantity" json:"uploads_quantity"`
	MembersQuantity uint       `db:"members_quantity" json:"members_quantity"`
	TagID           *string    `db:"tag_id" json:"tag_id"`
	CreatorID       string     `db:"creator_id" json:"creator_id"`
	DeletedBy       *string    `db:"deleted_by" json:"deleted_by"`
	CreatedAt       time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt       *time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt       *time.Time `db:"deleted_at" json:"deleted_at"`
}

func (c *Collection) Validate() error {
	if validation.IsEmpty(c.Name) {
		return &validation.ValidationError{
			Message: validation.NewRequiredFieldErrorMessage("name"),
		}
	}

	if validation.IsGreater(c.Name, 20) {
		return &validation.ValidationError{
			Message: validation.NewFieldLengthErrorMessage("name", 20, false),
		}
	}

	if validation.IsLower(c.Name, 2) {
		return &validation.ValidationError{
			Message: validation.NewFieldLengthErrorMessage("name", 2, true),
		}
	}

	if validation.IsEmpty(c.Secret) {
		return &validation.ValidationError{
			Message: validation.NewRequiredFieldErrorMessage("secret"),
		}
	}

	if validation.IsGreater(c.Secret, 20) {
		return &validation.ValidationError{
			Message: validation.NewFieldLengthErrorMessage("secret", 20, false),
		}
	}

	if validation.IsLower(c.Secret, 10) {
		return &validation.ValidationError{
			Message: validation.NewFieldLengthErrorMessage("secret", 10, true),
		}
	}

	if validation.IsEmpty(c.CreatorID) {
		return &validation.ValidationError{
			Message: validation.NewRequiredFieldErrorMessage("creator_id"),
		}
	}

	if validation.IsEmpty(c.Tag) {
		return &validation.ValidationError{
			Message: validation.NewRequiredFieldErrorMessage("tag"),
		}
	}

	return nil
}

func (c *Collection) encryptSecret() error {
	if validation.IsEmpty(c.Secret) {
		return &validation.ValidationError{
			Message: validation.NewRequiredFieldErrorMessage("secret"),
		}
	}

	if validation.IsLower(c.Secret, 8) {
		return &validation.ValidationError{
			Message: validation.NewFieldLengthErrorMessage("secret", 8, true),
		}
	}

	if validation.IsGreater(c.Secret, 16) {
		return &validation.ValidationError{
			Message: validation.NewFieldLengthErrorMessage("secret", 16, false),
		}
	}

	secretHash, err := cryptography.GenerateHash(c.Secret)
	if err != nil {
		return err
	}

	c.Secret = secretHash

	return nil
}
