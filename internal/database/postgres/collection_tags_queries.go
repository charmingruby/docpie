package postgres

const (
	createCollectionTag     = "create collection tag"
	findCollectionTagByName = "find collection tag by name"
	findCollectionTagByID   = "find collection tag by id"
)

func collectionTagsQueries() map[string]string {
	return map[string]string{
		createCollectionTag:     "INSERT INTO collection_tags (id, name, description) VALUES ($1, $2, $3) RETURNING *",
		findCollectionTagByName: "SELECT * FROM collection_tags WHERE name = $1",
		findCollectionTagByID:   "SELECT * FROM collection_tags WHERE id = $1",
	}
}
