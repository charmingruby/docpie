package collections

import (
	"github.com/charmingruby/upl/internal/domain"
	"github.com/charmingruby/upl/internal/domain/accounts"
	"github.com/charmingruby/upl/internal/validation/errs"
)

type UploadService struct {
	repo            UploadsRepository
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
		repo:            repo,
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

	if err := s.repo.Create(upload); err != nil {
		return err
	}

	collection, _ := s.collectionsRepo.FindByID(upload.CollectionID)
	collection.UploadsQuantity += 1
	if err := s.collectionsRepo.Save(&collection); err != nil {
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

func (s *UploadService) FetchCollectionUploads(page int, collectionID string) ([]Upload, *Collection, error) {
	collection, err := s.collectionsRepo.FindByID(collectionID)
	if err != nil {
		return nil, nil, err
	}

	uploads, err := s.repo.FetchUploadsByCollectionID(page, collectionID)
	if err != nil {
		return nil, nil, err
	}

	return uploads, &collection, nil
}
