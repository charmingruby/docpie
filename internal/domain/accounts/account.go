package accounts

import (
	"fmt"
	"time"

	"github.com/charmingruby/upl/internal/core"
	"github.com/charmingruby/upl/internal/validation"
	"github.com/charmingruby/upl/internal/validation/errs"
	"github.com/charmingruby/upl/pkg/cryptography"
)

const (
	managerRole = "manager"
	defaultRole = "member"
)

func NewAccount(name, lastName, email, password string) (*Account, error) {
	a := &Account{
		ID:                         core.NewId(),
		Name:                       name,
		LastName:                   lastName,
		Email:                      email,
		AvatarURL:                  nil,
		Password:                   password,
		CollectionsCreatedQuantity: 0,
		CollectionsMemberQuantity:  0,
		UploadQuantity:             0,
		DeletedBy:                  nil,
		CreatedAt:                  time.Now(),
		UpdatedAt:                  nil,
		DeletedAt:                  nil,
	}

	a.Role = a.accountRoles()[defaultRole]

	if err := a.Validate(); err != nil {
		return nil, err
	}

	if err := a.encryptPassword(); err != nil {
		return nil, err
	}

	return a, nil
}

type Account struct {
	ID                         string     `db:"id" json:"id"`
	Name                       string     `db:"name" json:"name"`
	LastName                   string     `db:"last_name" json:"last_name"`
	Email                      string     `db:"email" json:"email"`
	Role                       string     `db:"role" json:"role"`
	AvatarURL                  *string    `db:"avatar_url" json:"avatar_url"`
	CollectionsCreatedQuantity int        `db:"collections_created_quantity" json:"collections_created_quantity"`
	CollectionsMemberQuantity  int        `db:"collections_member_quantity" json:"collections_member_quantity"`
	UploadQuantity             int        `db:"upload_quantity" json:"upload_quantity"`
	Password                   string     `db:"password" json:"password"`
	DeletedBy                  *string    `db:"deleted_by" json:"deleted_by"`
	CreatedAt                  time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt                  *time.Time `db:"updated_at" json:"updated_at"`
	DeletedAt                  *time.Time `db:"deleted_at" json:"deleted_at"`
}

func (a *Account) Validate() error {
	if validation.IsEmpty(a.Name) {
		return &errs.ValidationError{
			Message: errs.EntitieisRequiredFieldErrorMessage("name"),
		}
	}

	if validation.IsLower(a.Name, 3) {
		return &errs.ValidationError{
			Message: errs.EntitiesFieldLengthErrorMessage("name", 3, true),
		}
	}

	if validation.IsGreater(a.Name, 16) {
		return &errs.ValidationError{
			Message: errs.EntitiesFieldLengthErrorMessage("name", 3, true),
		}
	}

	if validation.IsEmpty(a.LastName) {
		return &errs.ValidationError{
			Message: errs.EntitieisRequiredFieldErrorMessage("last name"),
		}
	}

	if validation.IsLower(a.LastName, 3) {
		return &errs.ValidationError{
			Message: errs.EntitiesFieldLengthErrorMessage("last name", 3, true),
		}
	}

	if validation.IsGreater(a.LastName, 32) {
		return &errs.ValidationError{
			Message: errs.EntitiesFieldLengthErrorMessage("last name", 32, false),
		}
	}

	if validation.IsEmpty(a.Email) {
		return &errs.ValidationError{
			Message: errs.EntitieisRequiredFieldErrorMessage("email"),
		}
	}

	if validation.IsLower(a.Email, 6) {
		return &errs.ValidationError{
			Message: errs.EntitiesFieldLengthErrorMessage("email", 6, true),
		}
	}

	if validation.IsGreater(a.Email, 64) {
		return &errs.ValidationError{
			Message: errs.EntitiesFieldLengthErrorMessage("email", 64, false),
		}
	}

	if !validation.IsEmail(a.Email) {
		return &errs.ValidationError{
			Message: "invalid email format",
		}
	}

	return nil
}

func (a *Account) Touch() {
	now := time.Now()
	a.UpdatedAt = &now
}

func (a *Account) SetRole(role string) {
	a.Role = role
}

func (a *Account) accountRoles() map[string]string {
	return map[string]string{
		managerRole: "manager",
		defaultRole: "member",
	}
}

func (a *Account) validateRole(role string) (string, error) {
	namedRole, ok := a.accountRoles()[role]

	if !ok {
		return "nil", fmt.Errorf("invalid role '%s'", role)
	}

	return namedRole, nil
}

func (a *Account) encryptPassword() error {
	if validation.IsEmpty(a.Password) {
		return &errs.ValidationError{
			Message: errs.EntitieisRequiredFieldErrorMessage("password"),
		}
	}

	if validation.IsLower(a.Password, 8) {
		return &errs.ValidationError{
			Message: errs.EntitiesFieldLengthErrorMessage("password", 8, true),
		}
	}

	if validation.IsGreater(a.Password, 16) {
		return &errs.ValidationError{
			Message: errs.EntitiesFieldLengthErrorMessage("password", 16, false),
		}
	}

	passwordHash, err := cryptography.GenerateHash(a.Password)
	if err != nil {
		return err
	}

	a.Password = passwordHash

	return nil
}
