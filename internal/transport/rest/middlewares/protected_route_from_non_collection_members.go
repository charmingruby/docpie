package middlewares

import (
	"fmt"
	"net/http"

	"github.com/charmingruby/upl/pkg/token"
	"github.com/gorilla/mux"
)

func (m *Middleware) ProtectedRouterFromNonNCollectionMembers(
	role string,
	next http.HandlerFunc,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		collectionID := params["id"]

		extractedToken := extractTokenFromRequest(r)
		if isTokenValid := token.NewJwtService().ValidateToken(extractedToken); !isTokenValid {
			m.logger.Error("Invalid token")
			sendResponse[any](w, "Unauthorized", http.StatusUnauthorized, nil)
			return
		}

		payload, err := token.NewJwtService().RetriveTokenPayload(extractedToken)
		if err != nil {
			m.logger.Error(err.Error())
			sendResponse[any](w, "Cannot retrieve token payload", http.StatusInternalServerError, nil)
			return
		}

		if _, err := m.collectionsRepository.FindByID(collectionID); err != nil {
			m.logger.Error("Collection not found")
			sendResponse[any](w, "Collection not found", http.StatusNotFound, nil)
			return
		}

		member, err := m.membersRepository.FindMemberInCollection(payload.AccountID, collectionID)
		if err != nil {
			m.logger.Error("Member not in collection")
			sendResponse[any](w, "Unauthorized, member not in collection", http.StatusUnauthorized, nil)
			return
		}

		requiresRole := len(role) != 0
		if requiresRole {
			allowed := role == member.Role

			if !allowed {
				msg := fmt.Sprintf("Member needs to be a %s", role)

				m.logger.Error(msg)
				sendResponse[any](w, fmt.Sprintf("Unauthorized, %s", msg), http.StatusUnauthorized, nil)
				return
			}
		}

		next(w, r)
	}
}
