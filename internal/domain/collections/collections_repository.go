package collections

type CollectionTagsRepository interface {
	FindByName(name string) (CollectionTag, error)
	Create(tag *CollectionTag) error
}

type CollectionsRepository interface {
	Create(collection *Collection) error
	FindByName(name string) (Collection, error)
}
