package accounts

type AccountRepository interface {
	Create(account *Account) error
}
