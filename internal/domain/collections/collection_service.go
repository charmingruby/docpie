package collections

import (
	"github.com/charmingruby/upl/internal/domain/accounts"
	"github.com/charmingruby/upl/internal/validation"
)

type CollectionService struct {
	repo         CollectionsRepository
	tagsRepo     CollectionTagsRepository
	accountsRepo accounts.AccountRepository
}

func NewCollectionService(repo CollectionsRepository, tagsRepo CollectionTagsRepository, accountsRepo accounts.AccountRepository) *CollectionService {
	return &CollectionService{
		repo:         repo,
		tagsRepo:     tagsRepo,
		accountsRepo: accountsRepo,
	}
}

func (s *CollectionService) Create(collection *Collection) error {
	_, err := s.repo.FindByName(collection.Name)
	if err == nil {
		return &validation.ServiceError{
			Message: validation.NewUniqueValidationErrorMessage(collection.Name),
		}
	}

	tag, err := s.tagsRepo.FindByName(collection.Tag)
	if err != nil {
		resourceNotFoundError := &validation.ResourceNotFoundError{
			Message: validation.NewResourceNotFoundErrorMessage("collection_tag"),
		}

		return resourceNotFoundError
	}

	collection.TagID = &tag.ID

	if err := s.repo.Create(collection); err != nil {
		return err
	}

	// Create a collection members as owner
	// Increments the collections members
	// Update account

	return nil
}
