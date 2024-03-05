package rest

import (
	"net/http"

	"github.com/charmingruby/upl/internal/domain/accounts"
	"github.com/charmingruby/upl/internal/transport/rest/endpoints"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type AccountHandler struct {
	accountService *accounts.AccountService
	logger         *logrus.Logger
}

func NewAccountHandler(logger *logrus.Logger, accountService *accounts.AccountService) *AccountHandler {
	return &AccountHandler{
		accountService: accountService,
		logger:         logger,
	}
}

func (h *AccountHandler) Register(router *mux.Router) {
	registerEndpoint := endpoints.MakeRegisterEndpoint(h.logger, h.accountService)
	authenticateEndpoint := endpoints.MakeAuthenticateEndpoint(h.logger, h.accountService)

	router.HandleFunc("/register", registerEndpoint).Methods(http.MethodPost)
	router.HandleFunc("/authenticate", authenticateEndpoint).Methods(http.MethodPost)
}
