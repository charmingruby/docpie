package postgres

import (
	"fmt"

	"github.com/charmingruby/upl/internal/domain/accounts"
	"github.com/charmingruby/upl/internal/validation"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

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
		return nil, fmt.Errorf("accounts repository wasn't able to build all the statements")
	}

	return &AccountsRepository{
		DB:         db,
		statements: sqlStmts,
	}, nil
}

func (r *AccountsRepository) statement(queryName string) (*sqlx.Stmt, error) {
	stmt, ok := r.statements[queryName]
	if !ok {
		return nil, fmt.Errorf("prepared statement '%s' not found", queryName)
	}

	return stmt, nil
}

func (r *AccountsRepository) Create(account *accounts.Account) error {
	stmt, err := r.statement(createAccount)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(account.ID, account.Name, account.LastName, account.Email, account.Role, account.Password, account.AvatarURL, account.UploadQuantity, account.CollectionsMemberQuantity, account.CollectionsCreatedQuantity)
	if err != nil {
		return fmt.Errorf("error %s %s: %v", "create", "account", err.Error())
	}

	return nil
}

func (r *AccountsRepository) FindByEmail(email string) (accounts.Account, error) {
	stmt, err := r.statement(findAccountByEmail)
	if err != nil {
		return accounts.Account{}, err
	}

	var a accounts.Account
	if err = stmt.Get(&a, email); err != nil {
		return accounts.Account{}, &validation.StorageError{
			Message: validation.NewResourceNotFoundByErrorMessage(email, "account", "email"),
		}
	}

	return a, nil
}

func (r *AccountsRepository) FindById(id string) (accounts.Account, error) {
	stmt, err := r.statement(findAccountById)
	if err != nil {
		return accounts.Account{}, err
	}

	var a accounts.Account
	if err := stmt.Get(&a, id); err != nil {
		return accounts.Account{}, &validation.StorageError{
			Message: validation.NewResourceNotFoundByErrorMessage(id, "account", "id"),
		}
	}

	return a, nil
}

func (r *AccountsRepository) Save(account *accounts.Account) error {
	stmt, err := r.statement(saveAccount)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		account.Name,
		account.LastName,
		account.Email,
		account.Role,
		account.AvatarURL,
		account.UpdatedAt,
		account.DeletedBy,
		account.DeletedAt,
		account.UploadQuantity,
		account.CollectionsMemberQuantity,
		account.CollectionsCreatedQuantity,
		account.ID,
	)

	if err != nil {
		return &validation.StorageError{
			Message: validation.NewQueryErrorMessage("account", "saving", err),
		}
	}

	return nil
}
