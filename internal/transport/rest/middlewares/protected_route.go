package middlewares

import (
	"net/http"

	"github.com/charmingruby/upl/pkg/token"
)

func (m *Middleware) ProtectedRoute(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		extractedToken := extractTokenFromRequest(r)

		if isTokenValid := token.NewJwtService().ValidateToken(extractedToken); !isTokenValid {
			m.logger.Error("Invalid token")
			sendResponse[any](w, "Unauthorized", http.StatusUnauthorized, nil)
			return
		}

		next(w, r)
	}
}
