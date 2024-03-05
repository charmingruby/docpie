package accounts

import (
	"github.com/charmingruby/upl/internal/validation"
)

type AccountService struct {
	AccountRepository AccountRepository
}

func NewAccountService(accountRepository AccountRepository) *AccountService {
	svc := &AccountService{AccountRepository: accountRepository}
	return svc
}

func (s *AccountService) Authenticate(email, password string) (*Account, error) {
	return nil, nil
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

func (s *AccountService) UploadAvatar() (*Account, error) {
	return nil, nil
}
