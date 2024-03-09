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
	ID           string     `json:"id"`
	Name         string     `json:"name"`
	Url          string     `json:"url"`
	FileSize     int64      `json:"file_size"`
	FileMimetype string     `json:"file_mimetype"`
	CollectionID string     `json:"collection_id"`
	UploaderID   string     `json:"uploader_id"`
	UploadedAt   time.Time  `json:"uploaded_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
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
