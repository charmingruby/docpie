package endpoints

import (
	"net/http"

	"github.com/charmingruby/upl/internal/domain/collections"
)

func MakeCreateCollectionEndpoint(collectionsService *collections.CollectionService, tagService *collections.CollectionTagService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
