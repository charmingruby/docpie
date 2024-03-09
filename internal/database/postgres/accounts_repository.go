package postgres

import (
	"fmt"

	"github.com/charmingruby/upl/internal/domain/accounts"
	"github.com/charmingruby/upl/internal/validation/errs"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type AccountsRepository struct {
	DB         *sqlx.DB
	statements map[string]*sqlx.Stmt
}

func NewAccountsRepository(logger *logrus.Logger, db *sqlx.DB) (*AccountsRepository, error) {
	sqlStmts := make(map[string]*sqlx.Stmt)

	var es []error
	for queryName, query := range accountQueries() {
		stmt, err := db.Preparex(query)
		if err != nil {
			msg := errs.DatabaseQueryPreparationErrorMessage(queryName, err.Error())
			logger.Error(msg)
			es = append(es, err)
		}

		sqlStmts[queryName] = stmt
	}

	if len(es) > 0 {
		return nil, &errs.DatabaseError{
			Message: errs.DatabaseRepositoryNotAbleErrorMessage("Accounts"),
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
		return nil, &errs.DatabaseError{
			Message: errs.DatabaseQueryNotPreparedErrorMessage(queryName),
		}
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
		return accounts.Account{}, &errs.DatabaseError{
			Message: errs.DatabaseResourceNotFoundErrorMessage("Account"),
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
		return accounts.Account{}, &errs.DatabaseError{
			Message: errs.DatabaseResourceNotFoundErrorMessage("Account"),
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
		return &errs.DatabaseError{
			Message: errs.DatabaseResourceNotFoundErrorMessage("Account"),
		}
	}

	return nil
}
