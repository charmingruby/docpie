package collections

import "github.com/charmingruby/upl/internal/validation/errs"

type CollectionTagService struct {
	repo CollectionTagsRepository
}

func NewCollectionTagsService(repo CollectionTagsRepository) *CollectionTagService {
	return &CollectionTagService{
		repo: repo,
	}
}

func (s *CollectionTagService) Create(tag *CollectionTag) error {
	_, err := s.repo.FindByName(tag.Name)
	if err == nil {
		return &errs.ServiceError{
			Message: errs.ServicesUniqueValidationErrorMessage("Name"),
		}
	}

	if err := s.repo.Create(tag); err != nil {
		return err
	}

	return nil
}
