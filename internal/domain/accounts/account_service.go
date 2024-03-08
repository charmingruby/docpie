package accounts

import (
	"time"

	"github.com/charmingruby/upl/internal/validation/errs"
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
		resourceNotFoundError := &errs.ResourceNotFoundError{
			Message: errs.ServicesInvalidCredentialsErrorMessage(),
		}

		return nil, resourceNotFoundError
	}

	isPasswordValid := cryptography.VerifyIfHashMatches(account.Password, password)
	if !isPasswordValid {
		credentialsNotMatchError := &errs.ServiceError{
			Message: errs.ServicesInvalidCredentialsErrorMessage(),
		}

		return nil, credentialsNotMatchError
	}

	return &account, nil
}

func (s *AccountService) Register(account *Account) error {
	_, err := s.repo.FindByEmail(account.Email)

	if err == nil {
		return &errs.ServiceError{
			Message: errs.ServicesUniqueValidationErrorMessage("Email"),
		}
	}

	if err := s.repo.Create(account); err != nil {
		return err
	}

	return nil
}

func (s *AccountService) UpdateAnAccountRole(accountID, role string) error {
	account, err := s.repo.FindById(accountID)
	if err != nil {
		resourceNotFoundError := &errs.ServiceError{
			Message: errs.ServicesResourceNotFoundErrorMessage("Account"),
		}

		return resourceNotFoundError
	}

	namedRole, err := account.validateRole(role)
	if err != nil {
		return err
	}

	if account.Role == namedRole {
		notModifiedError := &errs.NotModifiedError{
			Message: errs.ServicesNotModifiedErrorMessage(),
		}

		return notModifiedError
	}

	account.SetRole(namedRole)
	account.Touch()

	if err := s.repo.Save(&account); err != nil {
		return err
	}

	return nil
}

func (s *AccountService) UploadAvatar(accountID, fileURL string) error {
	account, err := s.repo.FindById(accountID)
	if err != nil {
		resourceNotFoundError := &errs.ServiceError{
			Message: errs.ServicesResourceNotFoundErrorMessage("Account"),
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
		resourceNotFoundError := &errs.ServiceError{
			Message: errs.ServicesResourceNotFoundErrorMessage("Account"),
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
