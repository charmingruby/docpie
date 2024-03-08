package endpoints

import (
	"fmt"
	"net/http"

	"github.com/charmingruby/upl/internal/domain/accounts"
	"github.com/charmingruby/upl/internal/validation"
	"github.com/charmingruby/upl/internal/validation/errs"
	"github.com/charmingruby/upl/pkg/token"
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
			payloadError := &errs.EndpointError{
				Message: errs.HTTPPayloadErrorMessage([]string{"email", "password"}),
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

			emptyPayloadFieldsError := &errs.EndpointError{
				Message: errs.HTTPEmptyPayloadFieldsErrorMessage(emptyFields),
			}

			logger.Error(emptyPayloadFieldsError)
			sendResponse[any](w, emptyPayloadFieldsError.Error(), http.StatusBadRequest, nil)
			return
		}

		account, err := accountsService.Authenticate(request.Email, request.Password)
		if err != nil {
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
		logger.Info(fmt.Sprintf("'%s' validated successfully credentials", request.Email))

		t, err := token.NewJwtService().GenerateToken(account.ID, account.Role)
		if err != nil {
			logger.Error(err)
			sendResponse[any](w, err.Error(), http.StatusBadRequest, nil)
			return
		}

		body := &AuthenticateResponse{
			Token: t,
		}

		msg := "Authenticated successfully."
		logger.Info(msg)
		sendResponse[AuthenticateResponse](w, msg, http.StatusOK, body)
	}
}
