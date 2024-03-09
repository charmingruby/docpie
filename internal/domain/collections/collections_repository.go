package collections

type CollectionsRepository interface {
	Create(collection *Collection) error
	FindByName(name string) (Collection, error)
	FindByID(id string) (Collection, error)
	Save(collection *Collection) error
}
