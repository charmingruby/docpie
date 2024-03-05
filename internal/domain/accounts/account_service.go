package domain

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

func (s *AccountService) Register(name, lastName, email, password string) error {
	account, err := NewAccount(name, lastName, email, password)
	if err != nil {
		return err
	}

	if err := s.AccountRepository.Create(account); err != nil {
		return err
	}

	return nil
}

func (s *AccountService) UploadAvatar() (*Account, error) {
	return nil, nil
}
