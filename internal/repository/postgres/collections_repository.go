package postgres

import (
	"github.com/charmingruby/upl/internal/domain/collections"
	"github.com/charmingruby/upl/internal/validation"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

const (
	createCollection     = "create collection"
	findCollectionByName = "find collection by name"
	saveCollection       = "save collection"
)

func collectionsQueries() map[string]string {
	return map[string]string{
		createCollection:     "INSERT INTO collections (id, name, description, secret, tag, tag_id, creator_id, deleted_by, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *",
		findCollectionByName: "SELECT * FROM collections WHERE name = $1",
		saveCollection:       "UPDATE collections SET name = $1, description = $2, tag = $3, tag_id = $4",
	}
}

type CollectionsRepository struct {
	DB         *sqlx.DB
	statements map[string]*sqlx.Stmt
}

func NewCollectionsRepository(logger *logrus.Logger, db *sqlx.DB) (*CollectionsRepository, error) {
	sqlStmts := make(map[string]*sqlx.Stmt)

	var errs []error
	for queryName, query := range collectionsQueries() {
		stmt, err := db.Preparex(query)
		if err != nil {
			logger.Errorf("error preparing statement %s: %v", queryName, err)
			errs = append(errs, err)
		}

		sqlStmts[queryName] = stmt
	}

	if len(errs) > 0 {
		return nil, &validation.StorageError{
			Message: validation.NewRepositoryStatementsPreparationErrorMessage("collections repository"),
		}
	}

	return &CollectionsRepository{
		DB:         db,
		statements: sqlStmts,
	}, nil
}

func (r *CollectionsRepository) statement(queryName string) (*sqlx.Stmt, error) {
	stmt, ok := r.statements[queryName]
	if !ok {
		return nil, &validation.StorageError{
			Message: validation.NewQueryStatementPreparationErrorMessage(queryName),
		}
	}

	return stmt, nil
}

func (r *CollectionsRepository) Create(collection *collections.Collection) error {
	stmt, err := r.statement(createCollection)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(collection.ID, collection.Name, collection.Description, collection.Secret, collection.Tag, collection.TagID, collection.CreatorID, collection.DeletedBy, collection.DeletedAt)
	if err != nil {
		return &validation.StorageError{
			Message: validation.NewQueryErrorMessage("collection", "creating", err),
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
		return collections.Collection{}, &validation.StorageError{
			Message: validation.NewResourceNotFoundByErrorMessage(name, "collection", "name"),
		}
	}

	return collection, nil
}
