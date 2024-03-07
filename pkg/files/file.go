package files

import (
	"fmt"
	"strings"

	"github.com/charmingruby/upl/internal/validation"
)

func NewFile(name, mimetype string, size int64, validMimetypes []string, maxSizeInBytes int64) (*File, error) {
	file := &File{
		Name:     name,
		Mimetype: mimetype,
		Size:     size,
	}

	if err := file.Validate(validMimetypes, maxSizeInBytes); err != nil {
		return nil, err
	}

	return file, nil
}

type File struct {
	Name     string `json:"name"`
	Mimetype string `json:"mimetype"`
	Size     int64  `json:"size"`
}

func (f *File) Validate(validMimetypes []string, maxSizeInBytes int64) error {
	var matchAMimetype bool

	for _, mimetype := range validMimetypes {
		if f.Mimetype == mimetype {
			matchAMimetype = true
		}
	}

	if !matchAMimetype {
		mimetypeError := &validation.FileError{
			Message: validation.NewInvalidMimetypeErrorMessage(f.Mimetype, validMimetypes),
		}

		return mimetypeError
	}

	if f.Size > maxSizeInBytes {
		sizeError := &validation.FileError{
			Message: validation.NewFileReachesMaximumSizeErrorMessage(f.Size, maxSizeInBytes),
		}

		return sizeError
	}

	return nil
}

func GetFileData(filename string) (string, string, error) {
	agg := strings.Split(filename, ".")

	if len(agg) < 2 {
		return "", "", fmt.Errorf("invalid file")
	}

	file := agg[0]
	mimetype := agg[1]

	return file, mimetype, nil
}
