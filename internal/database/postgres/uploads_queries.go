package postgres

const (
	createUpload = "create upload"
)

func uploadsQueries() map[string]string {
	return map[string]string{
		createUpload: `INSERT INTO collection_members (id, role, uploads_quantity, account_id, collection_id, left_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING *`,
	}
}
