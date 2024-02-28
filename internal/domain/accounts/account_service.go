package domain

import (
	"database/sql"
)

type AccountService struct {
	DB                *sql.DB
	AccountRepository AccountRepository
}

func NewAccountService(DB *sql.DB, accountRepository AccountRepository) *AccountService {
	svc := &AccountService{DB, accountRepository}
	return svc
}

func (a *AccountService) Authenticate(email, password string) (*Account, error) {

	return nil, nil
}

func (a *AccountService) Register(name, lastName, email, password string) (*Account, error) {
	return nil, nil
}

func (a *AccountService) UploadAvatar() (*Account, error) {
	return nil, nil
}
