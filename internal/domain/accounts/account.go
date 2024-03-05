package accounts

import (
	"time"

	"github.com/charmingruby/upl/internal/core"
	"github.com/charmingruby/upl/internal/validation"
	"github.com/charmingruby/upl/pkg/cryptography"
)

const (
	managerRole = "manager"
	defaultRole = "user"
)

func NewAccount(name, lastName, email, password string) (*Account, error) {
	a := &Account{
		ID:        core.NewId(),
		Name:      name,
		LastName:  lastName,
		Email:     email,
		Role:      defaultRole,
		AvatarURL: "",
		Password:  password,
	}

	if err := a.Validate(); err != nil {
		return nil, err
	}

	if err := a.encryptPassword(); err != nil {
		return nil, err
	}

	return a, nil
}

type Account struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	AvatarURL string    `json:"avatar_url"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
