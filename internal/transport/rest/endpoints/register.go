package endpoints

import (
	"fmt"
	"net/http"

	"github.com/charmingruby/upl/internal/domain/accounts"
	"github.com/charmingruby/upl/internal/validation"
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
			payloadError := &validation.EndpointError{
				Message: validation.NewPayloadErrorMessage([]string{"name", "last_name", "email", "password"}),
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

		msg := NewCreateResponse("Account")
		logger.Info(fmt.Sprintf("Account: '%s' created successfully.", newAccount.Email))
		sendResponse[any](
			w,
			msg,
			http.StatusCreated,
			nil,
		)
	}
}
