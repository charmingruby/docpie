package errs

type FileError struct {
	Message string `json:"message"`
}

func (e *FileError) Error() string {
	return e.Message
}
