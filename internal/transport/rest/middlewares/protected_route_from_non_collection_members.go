package middlewares

import (
	"fmt"
	"net/http"

	"github.com/charmingruby/upl/internal/domain/collections"
	"github.com/charmingruby/upl/pkg/token"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func ProtectedRouterFromNonNCollectionMembers(
	logger *logrus.Logger,
	repo collections.CollectionMembersRepository,
	next http.HandlerFunc,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		collectionID := params["id"]

		extractedToken := extractTokenFromRequest(r)
		if isTokenValid := token.NewJwtService().ValidateToken(extractedToken); !isTokenValid {
			logger.Error("Invalid token.")
			sendResponse[any](w, "Unauthorized.", http.StatusUnauthorized, nil)
			return
		}

		payload, err := token.NewJwtService().RetriveTokenPayload(extractedToken)
		if err != nil {
			logger.Error(err.Error())
			sendResponse[any](w, "Cannot retrieve token payload.", http.StatusInternalServerError, nil)
			return
		}

		member, err := repo.FindMemberInCollection(payload.AccountID, collectionID)
		if err != nil {
			logger.Error("Member not in collection.")
			sendResponse[any](w, "Unauthorized, member not in collection.", http.StatusUnauthorized, nil)
			return
		}

		fmt.Printf("%v\n", member)

		next(w, r)
	}
}
