package endpoints

import (
	"fmt"
	"net/http"
	"time"

	"github.com/charmingruby/upl/internal/domain/collections"
	"github.com/charmingruby/upl/pkg/cloudflare"
	"github.com/charmingruby/upl/pkg/files"
	"github.com/charmingruby/upl/pkg/token"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func MakeCreateUploadEndpoint(logger *logrus.Logger, uploadService *collections.UploadService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		collectionID := params["id"]

		extractedToken := extractTokenFromRequest(r)
		payload, err := token.NewJwtService().RetriveTokenPayload(extractedToken)
		if err != nil {
			logger.Error(err.Error())
			sendResponse[any](w, "Cannot retrieve token payload.", http.StatusInternalServerError, nil)
			return
		}

		maxUploadSize := files.MBToBytes(25)
		validMimetypes := []string{"jpg", "png", "jpeg", "pdf", "doc", "gif"}
		file, entity, err := handleMultipartFormFile(
			r,
			"upload",
			32,
			int64(maxUploadSize),
			validMimetypes,
		)
		if err != nil {
			logger.Error(err.Error())
			sendResponse[any](w, err.Error(), http.StatusBadRequest, nil)
			return

		}

		if err := entity.Validate(validMimetypes, int64(maxUploadSize)); err != nil {
			logger.Error(err.Error())
			sendResponse[any](w, "Cannot retrieve token payload.", http.StatusInternalServerError, nil)
			return
		}

		upload, err := collections.NewUpload(collectionID, payload.AccountID, entity)
		if err != nil {
			logger.Error(err.Error())
			sendResponse[any](w, err.Error(), http.StatusInternalServerError, nil)
			return

		}

		if err := uploadService.CreateUpload(upload); err != nil {
			logger.Error(err.Error())
			sendResponse[any](w, err.Error(), http.StatusBadRequest, nil)
			return
		}

		cl := cloudflare.New(logger)
		fileURL := fmt.Sprintf("%s-%d.%s", payload.AccountID, time.Now().Unix(), entity.Mimetype)
		if err := cl.UploadToBucket(file, fileURL); err != nil {
			logger.Error(err)
			sendResponse[any](w, "Unable to upload to Cloudflare", http.StatusInternalServerError, nil)
			return
		}

		msg := CreatedResponse("Upload")
		logger.Info(msg)
		sendResponse[any](w, msg, http.StatusCreated, nil)
	}
}
