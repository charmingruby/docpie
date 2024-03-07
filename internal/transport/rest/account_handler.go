package rest

import (
	"net/http"

	"github.com/charmingruby/upl/internal/domain/accounts"
	"github.com/charmingruby/upl/internal/transport/rest/endpoints"
	"github.com/charmingruby/upl/internal/transport/rest/middlewares"
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
	updateAnAccountRoleEndpoint := endpoints.MakeUpdateAnAccountRole(h.logger, h.accountService)
	uploadAvatarEndpoint := endpoints.MakeUploadAvatar(h.logger, h.accountService)
	deleteAnAccountEndpoint := endpoints.MakeDeleteAnAccountEndpoint(h.logger, h.accountService)

	router.HandleFunc("/register", registerEndpoint).Methods(http.MethodPost)
	router.HandleFunc("/authenticate", authenticateEndpoint).Methods(http.MethodPost)
	router.HandleFunc("/accounts/{id}/roles", middlewares.ProtectedRouteByRole(h.logger, "manager", updateAnAccountRoleEndpoint)).Methods(http.MethodPatch)
	router.HandleFunc("/me/avatar", middlewares.ProtectedRoute(h.logger, uploadAvatarEndpoint)).Methods(http.MethodPatch)
	router.HandleFunc("/accounts/{id}", middlewares.ProtectedRouteByRole(h.logger, "manager", deleteAnAccountEndpoint)).Methods(http.MethodDelete)
}
