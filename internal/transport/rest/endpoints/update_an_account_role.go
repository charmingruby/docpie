package endpoints

import (
	"fmt"
	"net/http"

	"github.com/charmingruby/upl/internal/domain/accounts"
	"github.com/charmingruby/upl/internal/validation"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type MakeUpdateAnAccountRoleRequest struct {
	Role string `json:"role"`
}

func MakeUpdateAnAccountRole(logger *logrus.Logger, accountsService *accounts.AccountService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		accountToUpdateId := vars["id"]

		request := &MakeUpdateAnAccountRoleRequest{}
		if err := parseRequest[MakeUpdateAnAccountRoleRequest](request, r.Body); err != nil {
			payloadError := &validation.EndpointError{Message: validation.NewPayloadErrorMessage([]string{"role"})}
			logger.Error(err.Error())
			sendResponse[any](w, payloadError.Error(), http.StatusBadRequest, nil)
			return
		}

		namedRole, err := accountsService.UpdateAnAccountRole(accountToUpdateId, request.Role)
		if err != nil {
			resourceNotFoundError, ok := err.(*validation.ResourceNotFoundError)
			if ok {
				logger.Error(resourceNotFoundError)
				sendResponse[any](w, resourceNotFoundError.Error(), http.StatusNotFound, nil)
				return
			}

			notModifiedError, ok := err.(*validation.NotModifiedError)
			if ok {
				logger.Error(notModifiedError)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusNotModified)
				return
			}

			logger.Error(err)
			sendResponse[any](w, err.Error(), http.StatusBadRequest, nil)
			return
		}

		msg := fmt.Sprintf("'%s' is now: '%s'", accountToUpdateId, namedRole)
		logger.Error(msg)
		sendResponse[any](w, msg, http.StatusOK, nil)
	}
}
