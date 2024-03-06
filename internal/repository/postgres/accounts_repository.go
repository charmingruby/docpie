package postgres

import (
	"github.com/charmingruby/upl/internal/domain/accounts"
	"github.com/charmingruby/upl/internal/validation"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

const (
	createAccount      = "create account"
	findAccountByEmail = "find by email"
	findAccountById    = "find by id"
	saveAccount        = "save account"
)

func accountQueries() map[string]string {
	return map[string]string{
		createAccount:      `INSERT INTO accounts (id, name, last_name, email, role, password, avatar_url) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING *`,
		findAccountByEmail: `SELECT * FROM accounts WHERE email = $1`,
		findAccountById:    `SELECT * FROM accounts WHERE id = $1`,
		saveAccount:        `UPDATE accounts SET name = $1, last_name = $2, email = $3, role = $4, avatar_url = $5 where id = $6`,
	}
}

type AccountsRepository struct {
	DB         *sqlx.DB
	statements map[string]*sqlx.Stmt
}

func NewAccountsRepository(logger *logrus.Logger, db *sqlx.DB) (*AccountsRepository, error) {
	sqlStmts := make(map[string]*sqlx.Stmt)

	var errs []error
	for queryName, query := range accountQueries() {
		stmt, err := db.Preparex(query)
		if err != nil {
			logger.Errorf("error preparing statement %s: %v", queryName, err)
			errs = append(errs, err)
		}

		sqlStmts[queryName] = stmt
	}

	if len(errs) > 0 {
		return nil, &validation.StorageError{
			Message: validation.NewRepositoryStatementsPreparationErrorMessage("accounts repository"),
		}
	}

	return &AccountsRepository{
		DB:         db,
		statements: sqlStmts,
	}, nil
}

func (r *AccountsRepository) statement(queryName string) (*sqlx.Stmt, error) {
	stmt, ok := r.statements[queryName]
	if !ok {
		return nil, &validation.StorageError{
			Message: validation.NewQueryStatementPreparationErrorMessage(queryName),
		}
	}

	return stmt, nil
}

func (r *AccountsRepository) Create(account *accounts.Account) error {
	stmt, err := r.statement(createAccount)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(account.ID, account.Name, account.LastName, account.Email, account.Role, account.Password, account.AvatarURL)
	if err != nil {
		return &validation.StorageError{
			Message: validation.NewQueryErrorMessage("account", "creating", err),
		}
	}

	return nil
}

func (r *AccountsRepository) FindByEmail(email string) (*accounts.Account, error) {
	stmt, err := r.statement(findAccountByEmail)
	if err != nil {
		return nil, err
	}

	account := accounts.Account{}
	if err = stmt.Get(&account, email); err != nil {
		return nil, &validation.StorageError{
			Message: validation.NewResourceNotFoundByErrorMessage(email, "account", "email"),
		}
	}

	return &account, nil
}

func (r *AccountsRepository) FindById(id string) (*accounts.Account, error) {
	stmt, err := r.statement(findAccountById)
	if err != nil {
		return nil, err
	}

	account := accounts.Account{}
	if err := stmt.Get(&account, id); err != nil {
		return nil, &validation.StorageError{
			Message: validation.NewResourceNotFoundByErrorMessage(id, "account", "id"),
		}
	}

	return &account, nil
}

func (r *AccountsRepository) Save(account *accounts.Account) error {
	stmt, err := r.statement(saveAccount)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(account.Name, account.LastName, account.Email, account.Role, account.AvatarURL, account.ID)
	if err != nil {
		return &validation.StorageError{
			Message: validation.NewQueryErrorMessage("account", "creating", err),
		}
	}

	return nil
}
