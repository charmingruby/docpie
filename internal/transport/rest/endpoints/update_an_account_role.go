package endpoints

import (
	"net/http"

	"github.com/charmingruby/upl/internal/domain/accounts"
	"github.com/charmingruby/upl/internal/validation/errs"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type MakeUpdateAnAccountRoleRequest struct {
	Role string `json:"role"`
}

func MakeUpdateAnAccountRole(logger *logrus.Logger, accountsService *accounts.AccountService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		accountToUpdateID := params["id"]

		request := &MakeUpdateAnAccountRoleRequest{}
		if err := parseRequest[MakeUpdateAnAccountRoleRequest](request, r.Body); err != nil {
			payloadError := &errs.EndpointError{Message: errs.HTTPPayloadErrorMessage([]string{"role"})}
			logger.Error(err.Error())
			sendResponse[any](w, payloadError.Error(), http.StatusBadRequest, nil)
			return
		}

		if err := accountsService.UpdateAnAccountRole(accountToUpdateID, request.Role); err != nil {
			resourceNotFoundError, ok := err.(*errs.ResourceNotFoundError)
			if ok {
				logger.Error(resourceNotFoundError.Error())
				sendResponse[any](w, resourceNotFoundError.Error(), http.StatusNotFound, nil)
				return
			}

			notModifiedError, ok := err.(*errs.NotModifiedError)
			if ok {
				logger.Error(notModifiedError)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotModified)
				return
			}

			logger.Error(err.Error())
			sendResponse[any](w, err.Error(), http.StatusBadRequest, nil)
			return
		}

		msg := ModifiedResponse("Account", "avatar")
		logger.Info(msg)
		sendResponse[any](w, msg, http.StatusOK, nil)
	}
}
