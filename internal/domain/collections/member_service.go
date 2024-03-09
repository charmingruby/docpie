package collections

import "github.com/charmingruby/upl/internal/domain/accounts"

type CollectionMembersService struct {
	repo         CollectionMembersRepository
	accountsRepo accounts.AccountRepository
}

func NewCollectionsMembersService(
	repo CollectionMembersRepository,
	accountsRepo accounts.AccountRepository,
) *CollectionMembersService {
	return &CollectionMembersService{
		repo:         repo,
		accountsRepo: accountsRepo,
	}
}

func (s *CollectionMembersService) CreateMember(accountID, collectionID string) error {
	member, err := NewCollectionMember(accountID, collectionID)
	if err != nil {
		return err
	}

	if err := s.repo.Create(member); err != nil {
		return err
	}

	return nil
}
