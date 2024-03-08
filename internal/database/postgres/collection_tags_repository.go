package postgres

import (
	"github.com/charmingruby/upl/internal/domain/collections"
	"github.com/charmingruby/upl/internal/validation"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type CollectionTagsRepository struct {
	DB         *sqlx.DB
	statements map[string]*sqlx.Stmt
}

func NewCollectionTagsRepository(logger *logrus.Logger, db *sqlx.DB) (*CollectionTagsRepository, error) {
	sqlStmts := make(map[string]*sqlx.Stmt)

	var errs []error
	for queryName, query := range collectionTagsQueries() {
		stmt, err := db.Preparex(query)
		if err != nil {
			logger.Errorf("error preparing statement %s: %v", queryName, err)
			errs = append(errs, err)
		}

		sqlStmts[queryName] = stmt
	}

	if len(errs) > 0 {
		return nil, &validation.StorageError{
			Message: validation.NewRepositoryStatementsPreparationErrorMessage("collection tags repository"),
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
		return nil, &validation.StorageError{
			Message: validation.NewQueryStatementPreparationErrorMessage(queryName),
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
		return &validation.StorageError{
			Message: validation.NewQueryErrorMessage("collection tag", "creating", err),
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
		return collections.CollectionTag{}, &validation.StorageError{
			Message: validation.NewResourceNotFoundByErrorMessage(name, "collection tag", "name"),
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

		return collections.CollectionTag{}, &validation.StorageError{
			Message: validation.NewResourceNotFoundByErrorMessage(id, "collection tag", "id"),
		}
	}

	return tag, nil
}
