package accounts

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockAccountRepository struct {
	mock.Mock
}

func makeAccount() *Account {
	account, _ := NewAccount(
		"john",
		"doe",
		"john@doe.com",
		"secret123",
	)
	return account
}

func (m *MockAccountRepository) Create(account *Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *MockAccountRepository) FindByEmail(email string) (*Account, error) {
	args := m.Called(email)
	if account := args.Get(0); account != nil {
		return account.(*Account), args.Error(1)
	}

	return nil, args.Error(1)
}

func TestCreateCategory(t *testing.T) {
	repo := new(MockAccountRepository)
	service := NewAccountService(repo)

	acc := makeAccount()

	repo.On("Create", acc).Return(nil)

	err := service.Register(acc)

	assert.NoError(t, err)

	repo.AssertExpectations(t)
}
