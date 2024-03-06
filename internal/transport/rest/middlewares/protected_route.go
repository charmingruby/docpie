package middlewares

import (
	"net/http"
	"strings"

	"github.com/charmingruby/upl/pkg/token"
)

func ProtectedRoute(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		extractedToken := extractTokenFromRequest(r)

		if isTokenValid := token.NewJwtService().ValidateToken(extractedToken); !isTokenValid {
			sendResponse[any](w, "Unauthorized.", http.StatusUnauthorized, nil)
			return
		}

		next(w, r)
	}
}

func extractTokenFromRequest(req *http.Request) string {
	token := req.Header.Get("Authorization")

	splittedToken := strings.Split(token, " ")

	if len(splittedToken) == 2 {
		return splittedToken[1]
	}

	return ""
}
