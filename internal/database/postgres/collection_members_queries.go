package postgres

const (
	createCollectionMember = "create collection member"
	findMemberInCollection = "find member in collection"
)

func collectionMembersQueries() map[string]string {
	return map[string]string{
		createCollectionMember: `INSERT INTO collection_members (id, role, uploads_quantity, account_id, collection_id, left_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING *`,
		findMemberInCollection: `SELECT * FROM collection_members WHERE collection_id = $1 AND account_id = $2`,
	}
}
