package endpoints

import (
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/charmingruby/upl/internal/validation/errs"
	"github.com/charmingruby/upl/pkg/files"
)

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Body    any    `json:"body,omitempty"`
}

func (r *Response) Marshal() []byte {
	jsonResponse, err := json.Marshal(r)
	if err != nil {
		log.Printf("Failed to marshal response: %v", err)
	}

	return jsonResponse
}

func sendResponse[T any](rw http.ResponseWriter, message string, code int, body *T) {
	response := &Response{
		Code:    code,
		Message: message,
		Body:    body,
	}

	writeResponse(rw, response)
}

func writeResponse(rw http.ResponseWriter, r *Response) {
	jsonResponse := r.Marshal()

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(r.Code)

	_, err := rw.Write(jsonResponse)
	if err != nil {
		log.Printf("Error writing response: %v", err)
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func parseRequest[T any](r *T, body io.ReadCloser) error {
	if err := json.NewDecoder(body).Decode(&r); err != nil {
		return err
	}

	return nil
}

func extractTokenFromRequest(req *http.Request) string {
	token := req.Header.Get("Authorization")

	splittedToken := strings.Split(token, " ")

	if len(splittedToken) == 2 {
		return splittedToken[1]
	}

	return ""
}

func handleMultipartFormFile(r *http.Request, key string, multiformMemory int64, fileMaxSizeInBytes int64, validMimetypes []string) (multipart.File, *files.File, error) {
	r.ParseMultipartForm(multiformMemory << 20)
	multipartFormKey := key
	file, fileHeader, err := r.FormFile(multipartFormKey)
	if err != nil {
		noFileFoundError := &errs.FileError{
			Message: errs.FilesNoFileErrorMessage(multipartFormKey),
		}

		return nil, nil, noFileFoundError
	}

	// Validate file
	filename, mimetype, err := files.GetFileData(fileHeader.Filename)
	if err != nil {
		return nil, nil, err

	}

	//validMimetypes := []string{"jpg", "png", "jpeg"}
	//maxSizeInBytes := 10000000 // 10 mb

	fileEntity, err := files.NewFile(filename, mimetype, fileHeader.Size, validMimetypes, int64(fileMaxSizeInBytes))
	if err != nil {
		return nil, nil, err
	}

	return file, fileEntity, nil
}
