package collections

type CollectionsRepository interface {
	Create(collection *Collection) error
	FindByName(name string) (Collection, error)
}
