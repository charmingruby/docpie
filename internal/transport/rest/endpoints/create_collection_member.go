package endpoints

import (
	"net/http"

	"github.com/charmingruby/upl/internal/domain/collections"
	"github.com/sirupsen/logrus"
)

func MakeCreateCollectionMemberEndpoint(logger *logrus.Logger, membersService *collections.CollectionMembersService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {}
}
