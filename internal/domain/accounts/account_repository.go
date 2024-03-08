package accounts

type AccountRepository interface {
	Create(account *Account) error
	FindByEmail(email string) (Account, error)
	FindById(id string) (Account, error)
	Save(account *Account) error
}
