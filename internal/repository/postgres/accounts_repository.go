package postgres

import (
	domain "github.com/charmingruby/upl/internal/domain/accounts"
	"github.com/charmingruby/upl/internal/validation"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

const (
	createAccount = "create account"
)

func accountQueries() map[string]string {
	return map[string]string{
		createAccount: `INSERT INTO accounts (name, last_name, email, password) VALUES($1, $2, $3, $4) RETURNING *`,
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

func (r *AccountsRepository) Create(account *domain.Account) error {
	stmt, err := r.statement(createAccount)
	if err != nil {
		return err
	}

	if err := stmt.Get(account.Name, account.LastName, account.Email, account.Password); err != nil {
		return &validation.StorageError{
			Message: validation.NewQueryError("account", "creating", err),
		}
	}

	return nil
}
