package collections

import (
	"fmt"
	"time"

	"github.com/charmingruby/upl/internal/core"
	"github.com/charmingruby/upl/internal/validation"
	"github.com/charmingruby/upl/internal/validation/errs"
	"github.com/charmingruby/upl/pkg/files"
)

func NewUpload(collectionID, uploaderID string, file *files.File) (*Upload, error) {
	url := fmt.Sprintf("%s-%d.%s", uploaderID, time.Now().Unix(), file.Mimetype)

	upload := Upload{
		ID:           core.NewId(),
		Name:         file.Name,
		Url:          url,
		FileSize:     file.Size,
		FileMimetype: file.Mimetype,
		CollectionID: collectionID,
		UploaderID:   uploaderID,
		UploadedAt:   time.Now(),
		DeletedAt:    nil,
	}

	if err := upload.Validate(); err != nil {
		return nil, err
	}

	return &upload, nil
}

type Upload struct {
	ID           string     `db:"id" json:"id"`
	Name         string     `db:"name" json:"name"`
	Url          string     `db:"url" json:"url"`
	FileSize     int64      `db:"file_size" json:"file_size"`
	FileMimetype string     `db:"file_mimetype" json:"file_mimetype"`
	CollectionID string     `db:"collection_id" json:"collection_id"`
	UploaderID   string     `db:"uploader_id" json:"uploader_id"`
	UploadedAt   time.Time  `db:"uploaded_at" json:"uploaded_at"`
	DeletedAt    *time.Time `db:"deleted_at" json:"deleted_at"`
}

func (u *Upload) Validate() error {
	if validation.IsEmpty(u.CollectionID) {
		return &errs.ValidationError{
			Message: errs.EntitieisRequiredFieldErrorMessage("collection_id"),
		}
	}

	if validation.IsEmpty(u.UploaderID) {
		return &errs.ValidationError{
			Message: errs.EntitieisRequiredFieldErrorMessage("uploader_id"),
		}
	}

	return nil
}
