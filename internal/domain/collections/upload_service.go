package collections

import "github.com/charmingruby/upl/internal/domain/accounts"

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

	return nil
}
