package collections

type CollectionMembersRepository interface {
	Create(member *CollectionMember) error
}
