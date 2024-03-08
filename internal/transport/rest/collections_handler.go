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
}

func NewCollectionsHandler(logger *logrus.Logger, collectionTagService *collections.CollectionTagService) *CollectionsHandler {
	return &CollectionsHandler{
		logger:               logger,
		collectionTagService: collectionTagService,
	}
}

func (h *CollectionsHandler) Register(router *mux.Router) {
	createCollectionTagEndpoint := endpoints.MakeCreateCollectionTagEndpoint(h.logger, h.collectionTagService)

	router.HandleFunc("/collections/tags", middlewares.ProtectedRouteByRole(h.logger, "manager", createCollectionTagEndpoint)).Methods(http.MethodPost)
}
