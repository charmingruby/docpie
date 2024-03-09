package collections

import (
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

	if owner.CollectionsCreatedQuantity > 3 {
		return &errs.ValidationError{
			Message: "Members can only create 3 collections",
		}
	}

	if owner.CollectionsMemberQuantity > 10 {
		return &errs.ValidationError{
			Message: "Members can only be member of 10 collections",
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

	member, err := NewCollectionMember("manager", owner.ID, collection.ID)
	if err != nil {
		return err
	}
	if err := s.membersRepo.Create(member); err != nil {
		return err
	}

	collection.Touch()
	collection.MembersQuantity += 1

	owner.CollectionsCreatedQuantity += 1
	if err := s.accountsRepo.Save(&owner); err != nil {
		return err
	}

	return nil
}
