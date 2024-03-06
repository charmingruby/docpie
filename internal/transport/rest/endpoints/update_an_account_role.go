package endpoints

import (
	"net/http"

	"github.com/charmingruby/upl/internal/domain/accounts"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type MakeUpdateAnAccountRoleRequest struct{}

func MakeUpdateAnAccountRole(logger *logrus.Logger, accountsService *accounts.AccountService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		accountToUpdate := vars["id"]

	}
}
