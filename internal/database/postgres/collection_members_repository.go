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
	logger     *logrus.Logger
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
		logger:     logger,
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

	_, err = stmt.Exec(member.ID, member.Role, member.UploadsQuantity, member.AccountID, member.CollectionID, member.LeftAt, member.UpdatedAt)
	if err != nil {
		r.logger.Error(err.Error())

		return &errs.DatabaseError{
			Message: errs.DatabaseQueryErrorMessage("collection member", "creating", err),
		}
	}

	return nil
}

func (r *CollectionMembersRepository) FindMemberInCollection(accountID, collectionID string) (collections.CollectionMember, error) {
	stmt, err := r.statement(findMemberInCollection)
	if err != nil {
		return collections.CollectionMember{}, err
	}

	var member collections.CollectionMember
	if err := stmt.Get(&member, collectionID, accountID); err != nil {
		r.logger.Error(err.Error())

		return collections.CollectionMember{}, &errs.DatabaseError{
			Message: errs.DatabaseResourceNotFoundErrorMessage("Collection Member"),
		}
	}

	return member, nil
}

func (r *CollectionMembersRepository) FetchByCollectionID(collectionID string) ([]collections.CollectionMember, error) {
	stmt, err := r.statement(fetchMembersByCollectionID)
	if err != nil {
		return nil, err
	}

	var members []collections.CollectionMember
	if err := stmt.Select(&members, collectionID); err != nil {
		r.logger.Error(err.Error())

		return nil, &errs.DatabaseError{
			Message: err.Error(),
		}
	}

	return members, nil
}

func (r *CollectionMembersRepository) Save(member *collections.CollectionMember) error {
	stmt, err := r.statement(saveCollectionMember)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(member.Role, member.UploadsQuantity, member.UpdatedAt, member.LeftAt, member.ID); err != nil {
		r.logger.Error(err.Error())

		return &errs.DatabaseError{
			Message: "Unable to save member",
		}
	}

	return nil
}
