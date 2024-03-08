package collections

import "github.com/charmingruby/upl/internal/validation"

type CollectionTagService struct {
	collectionTagsRepository CollectionTagsRepository
}

func NewCollectionTagsService(collectionTagsRepository CollectionTagsRepository) *CollectionTagService {
	return &CollectionTagService{
		collectionTagsRepository: collectionTagsRepository,
	}
}

func (s *CollectionTagService) Create(tag *CollectionTag) error {
	isNameAvailable, _ := s.collectionTagsRepository.FindByName(tag.Name)
	if isNameAvailable != nil {
		return &validation.ServiceError{
			Message: validation.NewUniqueValidationErrorMessage(tag.Name),
		}
	}

	if err := s.collectionTagsRepository.Create(tag); err != nil {
		return err
	}

	return nil
}
