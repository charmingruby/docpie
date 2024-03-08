package postgres

import (
	"github.com/charmingruby/upl/internal/domain/collections"
	"github.com/charmingruby/upl/internal/validation"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type CollectionMembersRepository struct {
	DB         *sqlx.DB
	statements map[string]*sqlx.Stmt
}

func NewCollectionMembersRepository(logger *logrus.Logger, db *sqlx.DB) (*CollectionMembersRepository, error) {
	sqlStmts := make(map[string]*sqlx.Stmt)

	var errs []error
	for queryName, query := range collectionMembersQueries() {
		stmt, err := db.Preparex(query)
		if err != nil {
			logger.Errorf("error preparing statement %s: %v", queryName, err)
			errs = append(errs, err)
		}

		sqlStmts[queryName] = stmt
	}

	if len(errs) > 0 {
		return nil, &validation.StorageError{
			Message: validation.NewRepositoryStatementsPreparationErrorMessage("collection members repository"),
		}
	}

	return &CollectionMembersRepository{
		DB:         db,
		statements: sqlStmts,
	}, nil
}

func (r *CollectionMembersRepository) statement(queryName string) (*sqlx.Stmt, error) {
	stmt, ok := r.statements[queryName]
	if !ok {
		return nil, &validation.StorageError{
			Message: validation.NewQueryStatementPreparationErrorMessage(queryName),
		}
	}

	return stmt, nil
}

func (r *CollectionMembersRepository) Create(member *collections.CollectionMember) error {
	stmt, err := r.statement(createCollectionMember)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(member.ID, member.Role, member.UploadQuantity, member.AccountID, member.CollectionID, member.LeftAt, member.UpdatedAt)
	if err != nil {
		return &validation.StorageError{
			Message: validation.NewQueryErrorMessage("collection member", "creating", err),
		}
	}

	return nil
}
