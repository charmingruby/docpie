package postgres

const (
	createCollection     = "create collection"
	findCollectionByName = "find collection by name"
	findCollectionByID   = "find collection by id"
	saveCollection       = "save collection"
)

func collectionsQueries() map[string]string {
	return map[string]string{
		createCollection:     "INSERT INTO collections (id, name, description, secret, tag, tag_id, uploads_quantity, members_quantity, creator_id, deleted_by, deleted_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11) RETURNING *",
		findCollectionByName: "SELECT * FROM collections WHERE name = $1",
		findCollectionByID:   "SELECT * FROM collections WHERE id = $1",
		saveCollection:       "UPDATE collections SET name = $1, description = $2, tag = $3, tag_id = $4, uploads_quantity = $5, members_quantity = $6, deleted_by = $7, updated_at = $8, deleted_at = $9 WHERE id = $10",
	}
}
