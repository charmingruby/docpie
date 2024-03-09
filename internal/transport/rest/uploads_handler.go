package rest

import (
	"net/http"

	"github.com/charmingruby/upl/internal/domain/collections"
	"github.com/charmingruby/upl/internal/transport/rest/endpoints"
	"github.com/charmingruby/upl/internal/transport/rest/middlewares"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type UploadsHandler struct {
	logger        *logrus.Logger
	uploadService *collections.UploadService
	membersRepo   collections.CollectionMembersRepository
}

func NewUploadsHandler(logger *logrus.Logger, uploadService *collections.UploadService, membersRepo collections.CollectionMembersRepository) *UploadsHandler {
	return &UploadsHandler{
		logger:        logger,
		uploadService: uploadService,
		membersRepo:   membersRepo,
	}
}

func (h *UploadsHandler) Register(router *mux.Router) {
	createUploadEndpoint := endpoints.MakeCreateUploadEndpoint(h.logger, h.uploadService)

	router.HandleFunc("/collections/{id}/upload", middlewares.ProtectedRouterFromNonNCollectionMembers(h.logger, h.membersRepo, createUploadEndpoint)).Methods(http.MethodPost)
}
