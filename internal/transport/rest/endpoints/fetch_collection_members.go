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

type FetchCollectionMembersResponse struct {
	Page           int                            `json:"page"`
	TotalMembers   int                            `json:"total_members"`
	MembersFetched int                            `json:"members_fetched"`
	Members        []collections.CollectionMember `json:"members"`
}

func MakeFetchCollectionMembersEndpoint(
	logger *logrus.Logger,
	membersService *collections.CollectionMembersService,
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

		members, collection, err := membersService.FetchCollectionMembers(page, collectionID)
		if err != nil {
			isResourceNotFoundError := errors.Is(err, &errs.DatabaseError{})
			if isResourceNotFoundError {
				sendResponse[any](w, err.Error(), http.StatusNotFound, nil)
				return
			}

			sendResponse[any](w, err.Error(), http.StatusBadRequest, nil)
			return
		}

		res := &FetchCollectionMembersResponse{
			Page:           page,
			TotalMembers:   int(collection.MembersQuantity),
			MembersFetched: len(members),
			Members:        members,
		}
		msg := fmt.Sprintf("Fetched %d of %d members on page %d", res.MembersFetched, res.TotalMembers, res.Page)
		sendResponse[FetchCollectionMembersResponse](w, msg, http.StatusOK, res)
	}
}
