package endpoints

import (
	"net/http"

	"github.com/charmingruby/upl/internal/domain/collections"
	"github.com/charmingruby/upl/internal/validation/errs"
	"github.com/sirupsen/logrus"
)

type CreateCollectionTagRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func MakeCreateCollectionTagEndpoint(logger *logrus.Logger, collectionTagsService *collections.CollectionTagService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := CreateCollectionTagRequest{}
		if err := parseRequest(&request, r.Body); err != nil {
			payloadError := &errs.EndpointError{
				Message: errs.HTTPPayloadErrorMessage([]string{"name", "description"}),
			}

			logger.Error(payloadError.Error())
			sendResponse[any](w, payloadError.Error(), http.StatusBadRequest, nil)
			return
		}

		newTag, err := collections.NewCollectionTag(request.Name, request.Description)
		if err != nil {
			logger.Error(err.Error())
			sendResponse[any](w, err.Error(), http.StatusBadRequest, nil)
			return
		}

		if err := collectionTagsService.Create(newTag); err != nil {
			logger.Error(err.Error())
			sendResponse[any](w, err.Error(), http.StatusBadRequest, nil)
			return
		}

		msg := CreatedResponse("Collection Tag")
		logger.Info(msg)
		sendResponse[any](w, msg, http.StatusCreated, nil)
	}
}
