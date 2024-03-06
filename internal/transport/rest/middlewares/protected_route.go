package middlewares

import (
	"net/http"

	"github.com/charmingruby/upl/pkg/token"
	"github.com/sirupsen/logrus"
)

func ProtectedRoute(logger *logrus.Logger, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		extractedToken := extractTokenFromRequest(r)

		if isTokenValid := token.NewJwtService().ValidateToken(extractedToken); !isTokenValid {
			logger.Error("Invalid token.")
			sendResponse[any](w, "Unauthorized.", http.StatusUnauthorized, nil)
			return
		}

		next(w, r)
	}
}
