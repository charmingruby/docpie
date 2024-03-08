package collections

type CollectionTagsRepository interface {
	// Fetch
	// Get
	FindByName(name string) (*CollectionTag, error)
	Create(tag *CollectionTag) error
	// Save
	// Delete
}
