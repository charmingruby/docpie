package domain

type AccountRepository interface {
	Create(account *Account) error
}
