package collections

import (
	"fmt"

	"github.com/charmingruby/upl/internal/domain"
	"github.com/charmingruby/upl/internal/domain/accounts"
	"github.com/charmingruby/upl/internal/validation/errs"
)

type CollectionService struct {
	repo         CollectionsRepository
	tagsRepo     CollectionTagsRepository
	membersRepo  CollectionMembersRepository
	accountsRepo accounts.AccountRepository
}

func NewCollectionService(repo CollectionsRepository, tagsRepo CollectionTagsRepository, membersRepo CollectionMembersRepository, accountsRepo accounts.AccountRepository) *CollectionService {
	return &CollectionService{
		repo:         repo,
		tagsRepo:     tagsRepo,
		accountsRepo: accountsRepo,
		membersRepo:  membersRepo,
	}
}

func (s *CollectionService) Create(collection *Collection) error {
	owner, err := s.accountsRepo.FindById(collection.CreatorID)
	if err != nil {
		resourceNotFoundError := &errs.ResourceNotFoundError{
			Message: errs.ServicesResourceNotFoundErrorMessage("Account"),
		}

		return resourceNotFoundError
	}

	if owner.CollectionsCreatedQuantity >= domain.MaxMemberAccountCreatedCollections {
		return &errs.ServiceError{
			Message: fmt.Sprintf("Members can only create %d collections", domain.MaxMemberAccountCreatedCollections),
		}
	}

	_, err = s.repo.FindByName(collection.Name)
	if err == nil {
		return &errs.ServiceError{
			Message: errs.ServicesUniqueValidationErrorMessage(collection.Name),
		}
	}

	tag, err := s.tagsRepo.FindByID(collection.TagID)

	if err != nil {
		resourceNotFoundError := &errs.ResourceNotFoundError{
			Message: errs.ServicesResourceNotFoundErrorMessage("Collection Tag"),
		}

		return resourceNotFoundError
	}

	collection.Tag = &tag.Name

	if err := s.repo.Create(collection); err != nil {
		return err
	}

	member, err := NewCollectionMember(owner.ID, collection.ID)
	if err != nil {
		return err
	}
	if err := s.membersRepo.Create(member); err != nil {
		return err
	}

	collection.Touch()
	collection.MembersQuantity += 1
	if err := s.repo.Save(collection); err != nil {
		return err
	}

	owner.CollectionsCreatedQuantity += 1
	if err := s.accountsRepo.Save(&owner); err != nil {
		return err
	}

	return nil
}
