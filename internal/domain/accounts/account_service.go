package accounts

import (
	"github.com/charmingruby/upl/internal/validation"
	"github.com/charmingruby/upl/pkg/cryptography"
)

type AccountService struct {
	AccountRepository AccountRepository
}

func NewAccountService(accountRepository AccountRepository) *AccountService {
	svc := &AccountService{AccountRepository: accountRepository}
	return svc
}

func (s *AccountService) Authenticate(email, password string) (*Account, error) {
	account, err := s.AccountRepository.FindByEmail(email)
	if err != nil {
		resourceNotFoundError := &validation.ResourceNotFoundError{
			Message: validation.NewResourceNotFoundErrorMessage("account"),
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

	return account, nil
}

func (s *AccountService) Register(account *Account) error {
	emailAvailable, _ := s.AccountRepository.FindByEmail(account.Email)
	if emailAvailable != nil {
		return &validation.ServiceError{
			Message: validation.NewUniqueValidationErrorMessage(account.Email),
		}
	}

	if err := s.AccountRepository.Create(account); err != nil {
		return err
	}

	return nil
}

func (s *AccountService) UpdateAnAccountRole(accountID, role string) (string, error) {
	account, err := s.AccountRepository.FindById(accountID)
	if err != nil {
		resourceNotFoundError := &validation.ServiceError{
			Message: validation.NewResourceNotFoundErrorMessage("account"),
		}

		return "", resourceNotFoundError
	}

	namedRole, err := account.isRoleValid(role)
	if err != nil {
		return "", err
	}

	if account.Role == namedRole {
		notModifiedError := &validation.NotModifiedError{
			Message: validation.NewNotModifiedErrorMessage(account.ID, namedRole),
		}

		return "", notModifiedError
	}

	account.Role = namedRole

	if err := s.AccountRepository.Save(account); err != nil {
		return "", err
	}

	return account.Role, nil
}

func (s *AccountService) UploadAvatar() (*Account, error) {
	return nil, nil
}
