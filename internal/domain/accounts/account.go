package accounts

import (
	"fmt"
	"time"

	"github.com/charmingruby/upl/internal/core"
	"github.com/charmingruby/upl/internal/validation"
	"github.com/charmingruby/upl/pkg/cryptography"
)

const (
	managerRole = "manager"
	defaultRole = "default"
)

func NewAccount(name, lastName, email, password string) (*Account, error) {
	a := &Account{
		ID:        core.NewId(),
		Name:      name,
		LastName:  lastName,
		Email:     email,
		AvatarURL: "",
		Password:  password,
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
	ID        string    `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	LastName  string    `db:"last_name" json:"last_name"`
	Email     string    `db:"email" json:"email"`
	Role      string    `db:"role" json:"role"`
	AvatarURL string    `db:"avatar_url" json:"avatar_url"`
	Password  string    `db:"password" json:"password"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func (a *Account) Validate() error {
	if validation.IsEmpty(a.Name) {
		return &validation.ValidationError{
			Message: validation.NewRequiredFieldErrorMessage("name"),
		}
	}

	if validation.IsLower(a.Name, 3) {
		return &validation.ValidationError{
			Message: validation.NewFieldLengthErrorMessage("name", 3, true),
		}
	}

	if validation.IsGreater(a.Name, 16) {
		return &validation.ValidationError{
			Message: validation.NewFieldLengthErrorMessage("name", 3, true),
		}
	}

	if validation.IsEmpty(a.LastName) {
		return &validation.ValidationError{
			Message: validation.NewRequiredFieldErrorMessage("last name"),
		}
	}

	if validation.IsLower(a.LastName, 3) {
		return &validation.ValidationError{
			Message: validation.NewFieldLengthErrorMessage("last name", 3, true),
		}
	}

	if validation.IsGreater(a.LastName, 32) {
		return &validation.ValidationError{
			Message: validation.NewFieldLengthErrorMessage("last name", 32, false),
		}
	}

	if validation.IsEmpty(a.Email) {
		return &validation.ValidationError{
			Message: validation.NewRequiredFieldErrorMessage("email"),
		}
	}

	if validation.IsLower(a.Email, 6) {
		return &validation.ValidationError{
			Message: validation.NewFieldLengthErrorMessage("email", 6, true),
		}
	}

	if validation.IsGreater(a.Email, 64) {
		return &validation.ValidationError{
			Message: validation.NewFieldLengthErrorMessage("email", 64, false),
		}
	}

	if !validation.IsEmail(a.Email) {
		return &validation.ValidationError{
			Message: "invalid email format",
		}
	}

	return nil
}

func (a *Account) accountRoles() map[string]string {
	return map[string]string{
		managerRole: "manager",
		defaultRole: "member",
	}
}

func (a *Account) isRoleValid(role string) (string, error) {
	namedRole, ok := a.accountRoles()[role]

	if !ok {
		return "nil", fmt.Errorf("invalid role '%s'", role)
	}

	return namedRole, nil
}

func (a *Account) encryptPassword() error {
	if validation.IsEmpty(a.Password) {
		return &validation.ValidationError{
			Message: validation.NewRequiredFieldErrorMessage("password"),
		}
	}

	if validation.IsLower(a.Password, 8) {
		return &validation.ValidationError{
			Message: validation.NewFieldLengthErrorMessage("password", 8, true),
		}
	}

	if validation.IsGreater(a.Password, 16) {
		return &validation.ValidationError{
			Message: validation.NewFieldLengthErrorMessage("password", 16, false),
		}
	}

	passwordHash, err := cryptography.GenerateHash(a.Password)
	if err != nil {
		return err
	}

	a.Password = passwordHash

	return nil
}
