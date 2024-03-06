package rest

import (
	"net/http"

	"github.com/charmingruby/upl/internal/transport/rest/endpoints"
	"github.com/gorilla/mux"
)

type PingHandler struct{}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

func (h *PingHandler) Register(router *mux.Router) {
	pingEndpoint := endpoints.MakePingEndpoint()

	router.HandleFunc("/ping", pingEndpoint).Methods(http.MethodGet)
}
