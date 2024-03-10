package collections

import (
	"github.com/charmingruby/upl/internal/domain"
	"github.com/charmingruby/upl/internal/domain/accounts"
	"github.com/charmingruby/upl/internal/validation/errs"
)

type UploadService struct {
	Repo            UploadsRepository
	collectionsRepo CollectionsRepository
	accountsRepo    accounts.AccountRepository
	membersRepo     CollectionMembersRepository
}

func NewUploadService(
	repo UploadsRepository,
	collectionsRepo CollectionsRepository,
	accountsRepo accounts.AccountRepository,
	membersRepo CollectionMembersRepository,
) *UploadService {
	return &UploadService{
		Repo:            repo,
		collectionsRepo: collectionsRepo,
		accountsRepo:    accountsRepo,
		membersRepo:     membersRepo,
	}
}

func (s *UploadService) CreateUpload(upload *Upload) error {
	account, err := s.accountsRepo.FindById(upload.UploaderID)
	if err != nil {
		return err
	}

	if account.UploadQuantity >= domain.MaxMemberAccountUploadQuantity {
		return &errs.ServiceError{
			Message: "Member reached uploads limit",
		}
	}

	member, err := s.membersRepo.FindMemberInCollection(upload.UploaderID, upload.CollectionID)
	if err != nil {
		return err
	}

	if err := s.Repo.Create(upload); err != nil {
		return err
	}

	member.UploadsQuantity += 1
	if err := s.membersRepo.Save(&member); err != nil {
		return err
	}

	account.UploadQuantity += 1
	if err := s.accountsRepo.Save(&account); err != nil {
		return err
	}

	return nil
}
