package endpoints

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/charmingruby/upl/internal/domain/collections"
	"github.com/charmingruby/upl/internal/validation/errs"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type FetchCollectionUploadsResponse struct {
	Page           int                  `json:"page"`
	TotalUploads   int                  `json:"total_uploads"`
	UploadsFetched int                  `json:"uploads_fetched"`
	Uploads        []collections.Upload `json:"uploads"`
}

func MakeFetchCollectionUploadsEndpoints(
	logger *logrus.Logger,
	uploadsService *collections.UploadService,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		collectionID := params["id"]

		queryParams := r.URL.Query()
		pageParam := queryParams.Get("page")
		page, err := strconv.Atoi(pageParam)
		if err != nil {
			sendResponse[any](w, fmt.Sprintf("Invalid page param: %s", err.Error()), http.StatusBadRequest, nil)
		}

		uploads, collection, err := uploadsService.FetchCollectionUploads(page, collectionID)
		if err != nil {
			isResourceNotFoundError := errors.Is(err, &errs.DatabaseError{})
			if isResourceNotFoundError {
				sendResponse[any](w, err.Error(), http.StatusNotFound, nil)
				return
			}

			sendResponse[any](w, err.Error(), http.StatusBadRequest, nil)
			return
		}

		res := &FetchCollectionUploadsResponse{
			Page:           page + 1,
			TotalUploads:   int(collection.MembersQuantity),
			UploadsFetched: len(uploads),
			Uploads:        uploads,
		}
		msg := fmt.Sprintf("Fetched %d of %d uploads on page %d", res.UploadsFetched, res.TotalUploads, res.Page)
		sendResponse[FetchCollectionUploadsResponse](w, msg, http.StatusOK, res)
	}
}
