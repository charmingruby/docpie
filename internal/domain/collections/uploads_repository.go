package collections

type UploadsRepository interface {
	Create(upload *Upload) error
}
