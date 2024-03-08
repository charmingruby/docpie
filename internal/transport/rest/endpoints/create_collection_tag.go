package endpoints

import (
	"net/http"

	"github.com/charmingruby/upl/internal/domain/collections"
	"github.com/charmingruby/upl/internal/validation"
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
			payloadError := &validation.EndpointError{
				Message: validation.NewPayloadErrorMessage([]string{"name", "description"}),
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

		sendResponse[any](w, NewCreateResponse(newTag.Name), http.StatusCreated, nil)
	}
}
