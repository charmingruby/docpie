package collections

type UploadsRepository interface {
	Create(upload *Upload) error
	FetchUploadsByCollectionID(page int, collectionID string) ([]Upload, error)
}
