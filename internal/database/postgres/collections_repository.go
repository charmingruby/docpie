package postgres

import (
	"github.com/charmingruby/upl/internal/domain/collections"
	"github.com/charmingruby/upl/internal/validation/errs"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type CollectionsRepository struct {
	DB         *sqlx.DB
	statements map[string]*sqlx.Stmt
	logger     *logrus.Logger
}

func NewCollectionsRepository(logger *logrus.Logger, db *sqlx.DB) (*CollectionsRepository, error) {
	sqlStmts := make(map[string]*sqlx.Stmt)

	var es []error
	for queryName, query := range collectionsQueries() {
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
			Message: errs.DatabaseRepositoryNotAbleErrorMessage("collections"),
		}
	}

	return &CollectionsRepository{
		DB:         db,
		statements: sqlStmts,
		logger:     logger,
	}, nil
}

func (r *CollectionsRepository) statement(queryName string) (*sqlx.Stmt, error) {
	stmt, ok := r.statements[queryName]
	if !ok {
		return nil, &errs.DatabaseError{
			Message: errs.DatabaseQueryNotPreparedErrorMessage(queryName),
		}
	}

	return stmt, nil
}

func (r *CollectionsRepository) Create(collection *collections.Collection) error {
	stmt, err := r.statement(createCollection)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(collection.ID, collection.Name, collection.Description, collection.Secret, collection.Tag, collection.TagID, collection.UploadsQuantity, collection.MembersQuantity, collection.CreatorID, collection.DeletedBy, collection.DeletedAt)
	if err != nil {
		return &errs.DatabaseError{
			Message: errs.DatabaseQueryErrorMessage("collection", "creating", err),
		}
	}

	return nil
}

func (r *CollectionsRepository) FindByName(name string) (collections.Collection, error) {
	stmt, err := r.statement(findCollectionByName)
	if err != nil {
		return collections.Collection{}, err
	}

	var collection collections.Collection
	if err = stmt.Get(&collection, name); err != nil {
		return collections.Collection{}, &errs.DatabaseError{
			Message: errs.DatabaseResourceNotFoundErrorMessage("Collection"),
		}
	}

	return collection, nil
}

func (r *CollectionsRepository) Save(collections *collections.Collection) error {
	stmt, err := r.statement(saveCollection)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		collections.Name,
		collections.Description,
		collections.Tag,
		collections.TagID,
		collections.UploadsQuantity,
		collections.MembersQuantity,
		collections.DeletedBy,
		collections.UpdatedAt,
		collections.DeletedAt,
		collections.ID,
	)

	if err != nil {
		return &errs.DatabaseError{
			Message: errs.DatabaseResourceNotFoundErrorMessage("Collection"),
		}
	}

	return nil
}
