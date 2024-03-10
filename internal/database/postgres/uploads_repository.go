package postgres

import (
	"github.com/charmingruby/upl/internal/domain/collections"
	"github.com/charmingruby/upl/internal/validation/errs"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type UploadsRepository struct {
	DB         *sqlx.DB
	statements map[string]*sqlx.Stmt
	logger     *logrus.Logger
}

func NewUploadsRepository(logger *logrus.Logger, db *sqlx.DB) (*UploadsRepository, error) {
	sqlStmts := make(map[string]*sqlx.Stmt)

	var es []error
	for queryName, query := range uploadsQueries() {
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
			Message: errs.DatabaseRepositoryNotAbleErrorMessage("Uploads"),
		}
	}

	return &UploadsRepository{
		DB:         db,
		statements: sqlStmts,
		logger:     logger,
	}, nil
}

func (r *UploadsRepository) statement(queryName string) (*sqlx.Stmt, error) {
	stmt, ok := r.statements[queryName]
	if !ok {
		return nil, &errs.DatabaseError{
			Message: errs.DatabaseQueryNotPreparedErrorMessage(queryName),
		}
	}

	return stmt, nil
}

func (r *UploadsRepository) Create(upload *collections.Upload) error {
	stmt, err := r.statement(createUpload)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(upload.ID, upload.Name, upload.Url, upload.FileSize, upload.FileMimetype, upload.CollectionID, upload.UploaderID)
	if err != nil {
		return &errs.DatabaseError{
			Message: errs.DatabaseQueryErrorMessage("upload", "creating", err),
		}
	}

	return nil
}
