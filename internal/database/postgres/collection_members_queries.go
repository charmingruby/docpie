package postgres

const (
	createCollectionMember     = "create collection member"
	findMemberInCollection     = "find member in collection"
	fetchMembersByCollectionID = "fetch members by collection ID"
	saveCollectionMember       = "save collection member"
)

func collectionMembersQueries() map[string]string {
	return map[string]string{
		createCollectionMember: `INSERT INTO collection_members (id, role, uploads_quantity, account_id, collection_id, left_at, updated_at) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING *`,
		findMemberInCollection: `SELECT * FROM collection_members WHERE collection_id = $1 AND account_id = $2`,
		fetchMembersByCollectionID: `
			SELECT * FROM collection_members  
			WHERE collection_id = $1
			ORDER BY joined_at DESC
			LIMIT $2 
			OFFSET $3
		`,
		saveCollectionMember: `UPDATE collection_members SET role = $1, uploads_quantity = $2, updated_at = $3, left_at = $4 WHERE id = $5`,
	}
}
