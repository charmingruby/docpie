package accounts

import (
	"time"

	"github.com/charmingruby/upl/internal/validation"
	"github.com/charmingruby/upl/pkg/cryptography"
)

type AccountService struct {
	repo AccountRepository
}

func NewAccountService(repo AccountRepository) *AccountService {
	svc := &AccountService{repo: repo}
	return svc
}

func (s *AccountService) Authenticate(email, password string) (*Account, error) {
	account, err := s.repo.FindByEmail(email)
	if err != nil {
		resourceNotFoundError := &validation.ResourceNotFoundError{
			Message: validation.NewInvalidCredentialsErrorMessage(),
		}

		return nil, resourceNotFoundError
	}

	isPasswordValid := cryptography.VerifyIfHashMatches(account.Password, password)
	if !isPasswordValid {
		credentialsNotMatchError := &validation.ServiceError{
			Message: validation.NewInvalidCredentialsErrorMessage(),
		}

		return nil, credentialsNotMatchError
	}

	return &account, nil
}

func (s *AccountService) Register(account *Account) error {
	_, err := s.repo.FindByEmail(account.Email)

	if err == nil {
		return &validation.ServiceError{
			Message: validation.NewUniqueValidationErrorMessage("Email"),
		}
	}

	if err := s.repo.Create(account); err != nil {
		return err
	}

	return nil
}

func (s *AccountService) UpdateAnAccountRole(accountID, role string) (*Account, error) {
	account, err := s.repo.FindById(accountID)
	if err != nil {
		resourceNotFoundError := &validation.ServiceError{
			Message: validation.NewResourceNotFoundErrorMessage("Account"),
		}

		return nil, resourceNotFoundError
	}

	namedRole, err := account.validateRole(role)
	if err != nil {
		return nil, err
	}

	if account.Role == namedRole {
		notModifiedError := &validation.NotModifiedError{
			Message: validation.NewNotModifiedErrorMessage(),
		}

		return nil, notModifiedError
	}

	account.SetRole(namedRole)
	account.Touch()

	if err := s.repo.Save(&account); err != nil {
		return nil, err
	}

	return &account, nil
}

func (s *AccountService) UploadAvatar(accountID, fileURL string) error {
	account, err := s.repo.FindById(accountID)
	if err != nil {
		resourceNotFoundError := &validation.ServiceError{
			Message: validation.NewResourceNotFoundErrorMessage("Account"),
		}

		return resourceNotFoundError
	}

	account.AvatarURL = &fileURL
	account.Touch()

	if err := s.repo.Save(&account); err != nil {
		return err
	}

	return nil
}

func (s *AccountService) DeleteAnAccount(accountID, managerID string) error {
	account, err := s.repo.FindById(accountID)
	if err != nil {
		resourceNotFoundError := &validation.ServiceError{
			Message: validation.NewResourceNotFoundErrorMessage("Account"),
		}

		return resourceNotFoundError
	}

	account.Touch()
	account.DeletedBy = &managerID
	now := time.Now()
	account.DeletedAt = &now

	if err := s.repo.Save(&account); err != nil {
		return err
	}

	return nil
}
