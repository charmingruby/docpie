package postgres

import (
	"github.com/charmingruby/upl/internal/domain/accounts"
	"github.com/charmingruby/upl/internal/validation"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

const (
	createAccount = "create account"
	findByEmail   = "find by email"
)

func accountQueries() map[string]string {
	return map[string]string{
		createAccount: `INSERT INTO accounts (name, last_name, email, role, password, avatar_url) VALUES($1, $2, $3, $4, $5, $6) RETURNING *`,
		findByEmail:   `SELECT * FROM accounts WHERE email = $1`,
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
	_, err = stmt.Exec(account.Name, account.LastName, account.Email, account.Role, account.Password, account.AvatarURL)
	if err != nil {
		return &validation.StorageError{
			Message: validation.NewQueryErrorMessage("account", "creating", err),
		}
	}

	return nil
}

func (r *AccountsRepository) FindByEmail(email string) (*accounts.Account, error) {
	stmt, err := r.statement(findByEmail)
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
