package endpoints

import (
	"fmt"
	"net/http"
	"time"

	"github.com/charmingruby/upl/internal/domain/accounts"
	"github.com/charmingruby/upl/internal/validation"
	"github.com/charmingruby/upl/pkg/cloudflare"
	"github.com/charmingruby/upl/pkg/files"
	"github.com/charmingruby/upl/pkg/token"
	"github.com/sirupsen/logrus"
)

func MakeUploadAvatar(logger *logrus.Logger, accountsService *accounts.AccountService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract payload from token
		extractedToken := extractTokenFromRequest(r)
		payload, err := token.NewJwtService().RetriveTokenPayload(extractedToken)
		if err != nil {
			logger.Error(err.Error())
			sendResponse[any](w, "Cannot retrieve token payload.", http.StatusInternalServerError, nil)
			return
		}

		// Receive file from multipart form
		r.ParseMultipartForm(32 << 20)
		multipartFormKey := "avatar"
		file, fileHeader, err := r.FormFile(multipartFormKey)
		if err != nil {
			noFileFoundError := &validation.FileError{
				Message: validation.NewNoFileErrorMessage(multipartFormKey),
			}

			logger.Error(noFileFoundError.Error())
			sendResponse[any](w, noFileFoundError.Error(), http.StatusBadRequest, nil)
			return
		}

		// Validate file
		filename, mimetype, err := files.GetFileData(fileHeader.Filename)
		if err != nil {
			logger.Error(err.Error())
			sendResponse[any](w, err.Error(), http.StatusBadRequest, nil)
			return
		}

		validMimetypes := []string{"jpg", "png", "jpeg"}
		maxSizeInBytes := 10000000 // 10 mb
		fileEntity, err := files.NewFile(filename, mimetype, fileHeader.Size, validMimetypes, int64(maxSizeInBytes))
		if err != nil {
			logger.Error(err.Error())
			sendResponse[any](w, err.Error(), http.StatusBadRequest, nil)
			return
		}

		fileURL := fmt.Sprintf("%s-%d.%s", payload.AccountID, time.Now().Unix(), fileEntity.Mimetype)

		// Register file on Bucket
		cl := cloudflare.New(logger)
		if err = cl.UploadToBucket(file, fileURL); err != nil {
			logger.Error(err)
		}

		// Update account
		if err := accountsService.UploadAvatar(payload.AccountID, fileURL); err != nil {
			// Remove file from Bucket
			if err := cl.RemoveFromBucket(fileURL); err != nil {
				logger.Error(err.Error())
				sendResponse[any](w, err.Error(), http.StatusInternalServerError, nil)
				return
			}

			resourceNotFoundError, ok := err.(*validation.ResourceNotFoundError)
			if ok {
				logger.Error(resourceNotFoundError)
				sendResponse[any](w, resourceNotFoundError.Error(), http.StatusNotFound, nil)
				return
			}

			logger.Error(err)
			sendResponse[any](w, err.Error(), http.StatusBadRequest, nil)
			return
		}

		msg := "Avatar uploaded successfully."
		logger.Info(msg)
		sendResponse[any](w, msg, http.StatusOK, nil)
	}
}
