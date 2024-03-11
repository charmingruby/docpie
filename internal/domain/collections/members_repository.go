package collections

type CollectionMembersRepository interface {
	FindMemberInCollection(accountID, collectionID string) (CollectionMember, error)
	FetchByCollectionID(page int, collectionID string) ([]CollectionMember, error)
	Create(member *CollectionMember) error
	Save(member *CollectionMember) error
}
