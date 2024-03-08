package postgres

import (
	"github.com/charmingruby/upl/internal/domain/collections"
	"github.com/charmingruby/upl/internal/validation/errs"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type CollectionTagsRepository struct {
	DB         *sqlx.DB
	statements map[string]*sqlx.Stmt
}

func NewCollectionTagsRepository(logger *logrus.Logger, db *sqlx.DB) (*CollectionTagsRepository, error) {
	sqlStmts := make(map[string]*sqlx.Stmt)

	var es []error
	for queryName, query := range collectionTagsQueries() {
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
			Message: errs.DatabaseRepositoryNotAbleErrorMessage("Collection Tags"),
		}
	}

	return &CollectionTagsRepository{
		DB:         db,
		statements: sqlStmts,
	}, nil
}

func (r *CollectionTagsRepository) statement(queryName string) (*sqlx.Stmt, error) {
	stmt, ok := r.statements[queryName]
	if !ok {
		return nil, &errs.DatabaseError{
			Message: errs.DatabaseQueryNotPreparedErrorMessage(queryName),
		}
	}

	return stmt, nil
}

func (r *CollectionTagsRepository) Create(tag *collections.CollectionTag) error {
	stmt, err := r.statement(createCollectionTag)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(tag.ID, tag.Name, tag.Description)
	if err != nil {
		return &errs.DatabaseError{
			Message: errs.DatabaseQueryErrorMessage("collection tag", "creating", err),
		}
	}

	return nil
}

func (r *CollectionTagsRepository) FindByName(name string) (collections.CollectionTag, error) {
	stmt, err := r.statement(findCollectionTagByName)
	if err != nil {
		return collections.CollectionTag{}, err
	}

	var tag collections.CollectionTag
	if err = stmt.Get(&tag, name); err != nil {
		return collections.CollectionTag{}, &errs.DatabaseError{
			Message: errs.DatabaseResourceNotFoundErrorMessage("Collection Tag"),
		}
	}

	return tag, nil
}

func (r *CollectionTagsRepository) FindByID(id string) (collections.CollectionTag, error) {
	stmt, err := r.statement(findCollectionTagByID)
	if err != nil {
		return collections.CollectionTag{}, err
	}

	var tag collections.CollectionTag
	if err = stmt.Get(&tag, id); err != nil {

		return collections.CollectionTag{}, &errs.DatabaseError{
			Message: errs.DatabaseResourceNotFoundErrorMessage("Collection Tag"),
		}
	}

	return tag, nil
}
