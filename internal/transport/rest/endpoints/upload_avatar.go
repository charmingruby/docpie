package endpoints

import (
	"fmt"
	"net/http"
	"time"

	"github.com/charmingruby/upl/internal/domain/accounts"
	"github.com/charmingruby/upl/internal/storage"
	"github.com/charmingruby/upl/internal/validation/errs"
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

		file, entity, err := handleMultipartFormFile(
			r,
			"avatar",
			32,
			int64(files.MBToBytes(10)),
			[]string{"jpg", "png", "jpeg"},
		)
		if err != nil {
			logger.Error(err.Error())
			sendResponse[any](w, err.Error(), http.StatusBadRequest, nil)
			return
		}

		fileURL := fmt.Sprintf("%s-%d.%s", payload.AccountID, time.Now().Unix(), entity.Mimetype)

		// Register file on Bucket
		cl := storage.New(logger)
		if err = cl.UploadToBucket(file, fileURL); err != nil {
			logger.Error(err)
			sendResponse[any](w, "Unable to update avatar on Cloudflare", http.StatusInternalServerError, nil)
			return
		}

		// Update account
		if err := accountsService.UploadAvatar(payload.AccountID, fileURL); err != nil {
			// Remove file from Bucket
			if err := cl.RemoveFromBucket(fileURL); err != nil {
				logger.Error(err.Error())
				sendResponse[any](w, err.Error(), http.StatusInternalServerError, nil)
				return
			}

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

		msg := ModifiedResponse("Account", "avatar")
		logger.Info(msg)
		sendResponse[any](w, msg, http.StatusOK, nil)
	}
}
