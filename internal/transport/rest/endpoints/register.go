package endpoints

import (
	"net/http"

	"github.com/charmingruby/upl/internal/domain/accounts"
	"github.com/charmingruby/upl/internal/validation/errs"
	"github.com/sirupsen/logrus"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func MakeRegisterEndpoint(logger *logrus.Logger, accountService *accounts.AccountService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := &RegisterRequest{}
		if err := parseRequest[RegisterRequest](request, r.Body); err != nil {
			payloadError := &errs.EndpointError{
				Message: errs.HTTPPayloadErrorMessage([]string{"name", "last_name", "email", "password"}),
			}

			logger.Error(payloadError.Error())
			sendResponse[any](w, payloadError.Error(), http.StatusBadRequest, nil)
			return
		}

		newAccount, err := accounts.NewAccount(request.Name, request.LastName, request.Email, request.Password)
		if err != nil {
			logger.Error(err.Error())
			sendResponse[any](w, err.Error(), http.StatusBadRequest, nil)
			return
		}

		if err := accountService.Register(newAccount); err != nil {
			logger.Error(err.Error())
			sendResponse[any](w, err.Error(), http.StatusBadRequest, nil)
			return
		}

		msg := CreatedResponse("Account")
		logger.Info(msg)
		sendResponse[any](
			w,
			msg,
			http.StatusCreated,
			nil,
		)
	}
}
