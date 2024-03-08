package endpoints

import (
	"net/http"

	"github.com/charmingruby/upl/internal/domain/collections"
	"github.com/charmingruby/upl/internal/validation"
	"github.com/charmingruby/upl/internal/validation/errs"
	"github.com/charmingruby/upl/pkg/token"
	"github.com/sirupsen/logrus"
)

type CreateCollectionRequest struct {
	Name        string `json:"name"`
	Secret      string `json:"secret"`
	Description string `json:"description"`
	TagID       string `json:"tag_id"`
}

func MakeCreateCollectionEndpoint(logger *logrus.Logger, collectionsService *collections.CollectionService, tagService *collections.CollectionTagService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := CreateCollectionRequest{}
		if err := parseRequest[CreateCollectionRequest](&request, r.Body); err != nil {
			payloadError := &errs.EndpointError{
				Message: errs.HTTPPayloadErrorMessage([]string{"name", "secret", "tag"}),
			}

			logger.Error(payloadError)
			sendResponse[any](w, payloadError.Error(), http.StatusBadRequest, nil)
			return
		}

		extractedToken := extractTokenFromRequest(r)
		payload, err := token.NewJwtService().RetriveTokenPayload(extractedToken)
		if err != nil {
			logger.Error(err.Error())
			sendResponse[any](w, "Cannot retrieve token payload.", http.StatusInternalServerError, nil)
			return
		}

		IsNameEmpty := validation.IsEmpty(request.Name)
		IsSecretEmpty := validation.IsEmpty(request.Secret)
		IsTagIDEmpty := validation.IsEmpty(request.TagID)

		if IsNameEmpty || IsSecretEmpty || IsTagIDEmpty {
			var emptyFields []string

			if IsNameEmpty {
				emptyFields = append(emptyFields, "name")
			}

			if IsSecretEmpty {
				emptyFields = append(emptyFields, "secret")
			}

			if IsTagIDEmpty {
				emptyFields = append(emptyFields, "tag_id")
			}

			emptyPayloadFieldsError := &errs.EndpointError{
				Message: errs.HTTPEmptyPayloadFieldsErrorMessage(emptyFields),
			}

			logger.Error(emptyPayloadFieldsError)
			sendResponse[any](w, emptyPayloadFieldsError.Error(), http.StatusBadRequest, nil)
			return
		}

		collection, err := collections.NewCollection(request.Name, request.Secret, request.Description, request.TagID, payload.AccountID)
		if err != nil {
			logger.Error(err.Error())
			sendResponse[any](w, err.Error(), http.StatusBadRequest, nil)
			return
		}

		if err := collectionsService.Create(collection); err != nil {
			logger.Error(err.Error())
			sendResponse[any](w, err.Error(), http.StatusBadRequest, nil)
			return
		}
		msg := CreatedResponse("Collection")
		logger.Info(msg)
		sendResponse[any](w, msg, http.StatusCreated, nil)
	}
}
