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
	logger               *logrus.Logger
	collectionTagService *collections.CollectionTagService
	collectionsService   *collections.CollectionService
}

func NewCollectionsHandler(logger *logrus.Logger, collectionService *collections.CollectionService, collectionTagService *collections.CollectionTagService) *CollectionsHandler {
	return &CollectionsHandler{
		logger:               logger,
		collectionTagService: collectionTagService,
		collectionsService:   collectionService,
	}
}

func (h *CollectionsHandler) Register(router *mux.Router) {
	createCollectionTagEndpoint := endpoints.MakeCreateCollectionTagEndpoint(h.logger, h.collectionTagService)
	createCollectionEndpoint := endpoints.MakeCreateCollectionEndpoint(h.logger, h.collectionsService, h.collectionTagService)

	router.HandleFunc("/collections/tags", middlewares.ProtectedRouteByRole(h.logger, "manager", createCollectionTagEndpoint)).Methods(http.MethodPost)
	router.HandleFunc("/collections", middlewares.ProtectedRoute(h.logger, createCollectionEndpoint)).Methods(http.MethodPost)
}
