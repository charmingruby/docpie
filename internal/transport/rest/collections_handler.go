package rest

import (
	"net/http"

	"github.com/charmingruby/upl/internal/domain/collections"
	"github.com/charmingruby/upl/internal/transport/rest/endpoints"
	"github.com/charmingruby/upl/internal/transport/rest/middlewares"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type CollectionsHandler struct {
	logger             *logrus.Logger
	mw                 *middlewares.Middleware
	tagsService        *collections.CollectionTagService
	collectionsService *collections.CollectionService
	membersService     *collections.CollectionMembersService
	uploadsService     *collections.UploadService
}

func NewCollectionsHandler(
	logger *logrus.Logger,
	mw *middlewares.Middleware,
	collectionService *collections.CollectionService,
	tagsService *collections.CollectionTagService,
	membersService *collections.CollectionMembersService,
	uploadsService *collections.UploadService,
) *CollectionsHandler {
	return &CollectionsHandler{
		logger:             logger,
		mw:                 mw,
		tagsService:        tagsService,
		collectionsService: collectionService,
		membersService:     membersService,
		uploadsService:     uploadsService,
	}
}

func (h *CollectionsHandler) Register(router *mux.Router) {
	createCollectionTagEndpoint := endpoints.MakeCreateCollectionTagEndpoint(h.logger, h.tagsService)
	createCollectionEndpoint := endpoints.MakeCreateCollectionEndpoint(h.logger, h.collectionsService, h.tagsService)
	createCollectionMemberEndpoint := endpoints.MakeCreateCollectionMemberEndpoint(h.logger, h.membersService)
	createUploadEndpoint := endpoints.MakeCreateUploadEndpoint(h.logger, h.uploadsService)

	// Manager
	router.HandleFunc("/collections/tags", h.mw.ProtectedRouteByRole(h.logger, "manager", createCollectionTagEndpoint)).
		Methods(http.MethodPost)

	// Members
	router.HandleFunc("/collections", h.mw.ProtectedRoute(h.logger, createCollectionEndpoint)).
		Methods(http.MethodPost)
	router.HandleFunc("/collections/{id}/members", createCollectionMemberEndpoint).
		Methods(http.MethodPost)
	router.HandleFunc("/collections/{id}/upload", createUploadEndpoint).
		Methods(http.MethodPost)

}
