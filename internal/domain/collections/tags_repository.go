package collections

type CollectionTagsRepository interface {
	FindByName(name string) (CollectionTag, error)
	FindByID(id string) (CollectionTag, error)
	Create(tag *CollectionTag) error
}
