package postgres

const (
	createCollectionMember = "create collection member"
)

func collectionMembersQueries() map[string]string {
	return map[string]string{
		createCollectionMember: `INSERT INTO collection_members (id, role, upload_quantity, account_id, collection_id, left_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING *`,
	}
}
