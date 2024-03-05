package collections

import "time"

type Upload struct {
	ID           string     `json:"id"`
	Name         string     `json:"name"`
	Url          string     `json:"url"`
	FileName     string     `json:"file_name"`
	FileSize     int64      `json:"file_size"`
	FileMimetype string     `json:"file_mimetype"`
	CollectionID string     `json:"collection_id"`
	UploaderID   string     `json:"uploader_id"`
	UploadedAt   time.Time  `json:"uploaded_at"`
	DeletedAt    *time.Time `json:"deleted_at"`
}
