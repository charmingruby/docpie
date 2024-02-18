package rest

import (
	"github.com/charmingruby/make-it-survey/internal/transport/rest/endpoints"
	"github.com/gorilla/mux"
)

type PingHandler struct{}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

func (h *PingHandler) Register(router *mux.Router) {
	pingEndpoint := endpoints.MakePingEndpoint()

	router.HandleFunc("/ping", pingEndpoint)
}
