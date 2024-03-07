package validation

import "fmt"

type FileError struct {
	Message string `json:"message"`
}

func (e *FileError) Error() string {
	return e.Message
}

func NewInvalidMimetypeErrorMessage(mimetypeUnmatched string, validMimetypes []string) string {
	var msg = fmt.Sprintf(".%s is not a valid mimetype, please provide a valid mimetype: ", mimetypeUnmatched)

	for idx, mimetype := range validMimetypes {
		if idx+1 == len(mimetypeUnmatched) {
			msg += fmt.Sprintf("or .%s.", mimetype)
			continue
		}

		msg += fmt.Sprintf(".%s, ", mimetype)
	}

	return msg
}

func NewFileReachesMaximumSizeErrorMessage(currentSize, maxSize int64) string {
	return fmt.Sprintf("%d perpasses the limit of %d bytes", currentSize, maxSize)
}

func NewNoFileErrorMessage(key string) string {
	return fmt.Sprintf("no file found for key '%s'", key)
}
