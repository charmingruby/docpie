package postgres

import (
	"github.com/charmingruby/upl/internal/domain/collections"
	"github.com/charmingruby/upl/internal/validation/errs"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type CollectionMembersRepository struct {
	DB         *sqlx.DB
	statements map[string]*sqlx.Stmt
}

func NewCollectionMembersRepository(logger *logrus.Logger, db *sqlx.DB) (*CollectionMembersRepository, error) {
	sqlStmts := make(map[string]*sqlx.Stmt)

	var es []error
	for queryName, query := range collectionMembersQueries() {
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
			Message: errs.DatabaseRepositoryNotAbleErrorMessage("Collection Members"),
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
		return nil, &errs.DatabaseError{
			Message: errs.DatabaseQueryNotPreparedErrorMessage(queryName),
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
		return &errs.DatabaseError{
			Message: errs.DatabaseQueryErrorMessage("collection member", "creating", err),
		}
	}

	return nil
}
