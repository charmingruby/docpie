package endpoints

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
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
