package middlewares

import (
	"fmt"
	"net/http"

	"github.com/charmingruby/upl/pkg/token"
	"github.com/sirupsen/logrus"
)

func ProtectedRouteByRole(logger *logrus.Logger, role string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		if payload.Role != role {
			logger.Error(fmt.Sprintf("'%s' is a '%s' not a '%s", payload.AccountID, payload.Role, role))
			sendResponse[any](w, "Unauthorized by role.", http.StatusUnauthorized, nil)
			return
		}

		next(w, r)
	}
}
