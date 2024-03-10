package middlewares

import (
	"fmt"
	"net/http"

	"github.com/charmingruby/upl/pkg/token"
)

func (m *Middleware) ProtectedRouteByRole(role string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		extractedToken := extractTokenFromRequest(r)

		if isTokenValid := token.NewJwtService().ValidateToken(extractedToken); !isTokenValid {
			m.logger.Error("Invalid token.")
			sendResponse[any](w, "Unauthorized.", http.StatusUnauthorized, nil)
			return
		}

		payload, err := token.NewJwtService().RetriveTokenPayload(extractedToken)
		if err != nil {
			m.logger.Error(err.Error())
			sendResponse[any](w, "Cannot retrieve token payload.", http.StatusInternalServerError, nil)
			return
		}

		if payload.Role != role {
			m.logger.Error(fmt.Sprintf("'%s' is a '%s' not a '%s", payload.AccountID, payload.Role, role))
			sendResponse[any](w, "Unauthorized by role.", http.StatusUnauthorized, nil)
			return
		}

		next(w, r)
	}
}
