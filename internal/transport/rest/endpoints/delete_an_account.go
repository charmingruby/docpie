package endpoints

import (
	"net/http"

	"github.com/charmingruby/upl/internal/domain/accounts"
	"github.com/charmingruby/upl/internal/validation/errs"
	"github.com/charmingruby/upl/pkg/token"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func MakeDeleteAnAccountEndpoint(logger *logrus.Logger, accountsService *accounts.AccountService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		accountToDeleteID := params["id"]

		extractedToken := extractTokenFromRequest(r)
		payload, err := token.NewJwtService().RetriveTokenPayload(extractedToken)
		if err != nil {
			logger.Error(err.Error())
			sendResponse[any](w, "Cannot retrieve token payload.", http.StatusInternalServerError, nil)
			return
		}

		if err := accountsService.DeleteAnAccount(accountToDeleteID, payload.AccountID); err != nil {
			resourceNotFoundError, ok := err.(*errs.ResourceNotFoundError)
			if ok {
				logger.Error(resourceNotFoundError)
				sendResponse[any](w, resourceNotFoundError.Error(), http.StatusNotFound, nil)
				return
			}

			logger.Error(err)
			sendResponse[any](w, err.Error(), http.StatusBadRequest, nil)
			return
		}

		msg := DeleteResponse("Account")
		logger.Info(msg)
		sendResponse[any](w, msg, http.StatusOK, nil)
	}
}
