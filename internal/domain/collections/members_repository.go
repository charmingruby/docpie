package collections

type CollectionMembersRepository interface {
	FindMemberInCollection(accountID, collectionID string) (CollectionMember, error)
	FetchByCollectionID(collectionID string) ([]CollectionMember, error)
	Create(member *CollectionMember) error
	Save(member *CollectionMember) error
}
