package endpoints

import (
	"fmt"
	"net/http"

	"github.com/charmingruby/upl/internal/domain/accounts"
	"github.com/charmingruby/upl/internal/validation"
	"github.com/sirupsen/logrus"
)

type AuthenticateRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthenticateResponse struct {
	Token string `json:"token"`
}

func MakeAuthenticateEndpoint(logger *logrus.Logger, accountsService *accounts.AccountService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := &AuthenticateRequest{}
		if err := parseRequest[AuthenticateRequest](request, r.Body); err != nil {
			payloadError := &validation.EndpointError{
				Message: validation.NewPayloadErrorResponse([]string{"email", "password"}),
			}

			logger.Error(payloadError)
			sendResponse[any](w, payloadError.Error(), http.StatusBadRequest, nil)
			return
		}

		isEmailEmpty := validation.IsEmpty(request.Email)
		isPasswordEmpty := validation.IsEmpty(request.Password)

		if isEmailEmpty || isPasswordEmpty {
			var emptyFields []string

			if isEmailEmpty {
				emptyFields = append(emptyFields, "email")
			}

			if isPasswordEmpty {
				emptyFields = append(emptyFields, "password")
			}

			emptyPayloadFieldsError := &validation.EndpointError{
				Message: validation.NewEmptyPayloadFieldsErrorMessage(emptyFields),
			}

			logger.Error(emptyPayloadFieldsError)
			sendResponse[any](w, emptyPayloadFieldsError.Error(), http.StatusBadRequest, nil)
			return
		}

		err := accountsService.Authenticate(request.Email, request.Password)
		if err != nil {
			logger.Error(err)
			sendResponse[any](w, err.Error(), http.StatusBadRequest, nil)
			return
		}

		logger.Info(fmt.Sprintf("'%s' validated successfully credentials", request.Email))

	}
}
