package accounts

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

	if err := s.AccountRepository.Create(account); err != nil {
		return err
	}

	return nil
}

func (s *AccountService) UploadAvatar() (*Account, error) {
	return nil, nil
}
