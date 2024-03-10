package collections

import (
	"fmt"

	"github.com/charmingruby/upl/internal/domain"
	"github.com/charmingruby/upl/internal/domain/accounts"
	"github.com/charmingruby/upl/internal/validation/errs"
)

type CollectionMembersService struct {
	repo            CollectionMembersRepository
	accountsRepo    accounts.AccountRepository
	collectionsRepo CollectionsRepository
}

func NewCollectionsMembersService(
	repo CollectionMembersRepository,
	accountsRepo accounts.AccountRepository,
	collectionsRepo CollectionsRepository,
) *CollectionMembersService {
	return &CollectionMembersService{
		repo:            repo,
		accountsRepo:    accountsRepo,
		collectionsRepo: collectionsRepo,
	}
}

func (s *CollectionMembersService) CreateMember(accountID, collectionID string) error {
	member, err := NewCollectionMember(accountID, collectionID)
	if err != nil {
		return err
	}

	// Verify if accounts can be a member
	newMemberAccount, err := s.accountsRepo.FindById(accountID)
	if err != nil {
		return err
	}

	if newMemberAccount.CollectionsMemberQuantity > domain.MaxMemberAccountCollections {
		return &errs.ServiceError{
			Message: fmt.Sprintf("Members can only be member of %d collections", domain.MaxMemberAccountCollections),
		}
	}

	// Verify if accepts members
	collection, err := s.collectionsRepo.FindByID(collectionID)
	if err != nil {
		return err
	}

	if collection.MembersQuantity >= domain.MaxCollectionMembers {
		return &errs.ServiceError{
			Message: fmt.Sprintf("Collections can only have %d members", domain.MaxCollectionMembers),
		}
	}

	// Verify if is alreay member
	if _, err := s.repo.FindMemberInCollection(accountID, collectionID); err == nil {
		return &errs.ServiceError{
			Message: "Account is already a member",
		}
	}

	// Updates
	collection.MembersQuantity += 1
	newMemberAccount.CollectionsMemberQuantity += 1

	if err := s.accountsRepo.Save(&newMemberAccount); err != nil {
		return err
	}

	if err := s.collectionsRepo.Save(&collection); err != nil {
		return err
	}

	// Create member
	if err := s.repo.Create(member); err != nil {
		return err
	}

	return nil
}
