package collections

import (
	"github.com/charmingruby/upl/internal/domain/accounts"
	"github.com/charmingruby/upl/internal/validation/errs"
)

type UploadService struct {
	Repo            UploadsRepository
	collectionsRepo CollectionsRepository
	accountsRepo    accounts.AccountRepository
}

func NewUploadService(
	repo UploadsRepository,
	collectionsRepo CollectionsRepository,
	accountsRepo accounts.AccountRepository,
) *UploadService {
	return &UploadService{
		Repo:            repo,
		collectionsRepo: collectionsRepo,
		accountsRepo:    accountsRepo,
	}
}

func (s *UploadService) CreateUpload(upload *Upload) error {
	member, err := s.accountsRepo.FindById(upload.UploaderID)
	if err != nil {
		return err
	}

	if member.UploadQuantity > 20 {
		return &errs.ServiceError{
			Message: "Member reached uploads limit",
		}
	}

	if err := s.Repo.Create(upload); err != nil {
		return err
	}

	return nil
}
