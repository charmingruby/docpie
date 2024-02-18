package endpoints

import "net/http"

type PingResponse struct {
	Pong bool `json:"ping"`
}

func MakePingEndpoint() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		body := &PingResponse{
			Pong: true,
		}

		sendResponse[PingResponse](rw, "Pinged on API", http.StatusOK, body)
	}
}
