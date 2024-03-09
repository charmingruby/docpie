package endpoints

import (
	"net/http"

	"github.com/charmingruby/upl/internal/domain/collections"
	"github.com/charmingruby/upl/internal/validation/errs"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type AddMemberToCollectionRequest struct {
	AccountID string `json:"account_id"`
}

func MakeAddMemberToACollectionEndpoint(logger *logrus.Logger, membersService *collections.CollectionMembersService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		collectionID := params["id"]

		request := &AddMemberToCollectionRequest{}
		if err := parseRequest[AddMemberToCollectionRequest](request, r.Body); err != nil {
			payloadError := &errs.EndpointError{
				Message: errs.HTTPPayloadErrorMessage([]string{"account_id"}),
			}

			logger.Error(payloadError.Error())
			sendResponse[any](w, payloadError.Error(), http.StatusBadRequest, nil)
			return
		}

		if err := membersService.CreateMember(request.AccountID, collectionID); err != nil {
			logger.Error(err.Error())
			sendResponse[any](w, err.Error(), http.StatusBadRequest, nil)
			return
		}

		msg := CreatedResponse("Collection Member")
		logger.Info(msg)
		sendResponse[any](w, msg, http.StatusCreated, nil)
	}
}
