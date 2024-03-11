package postgres

const (
	createUpload               = "create upload"
	fetchUploadsByCollectionID = "fetch uploads by collection id"
)

func uploadsQueries() map[string]string {
	return map[string]string{
		createUpload: `INSERT INTO uploads (id, name, url, file_size, file_mimetype, collection_id, uploader_id) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING *`,
		fetchUploadsByCollectionID: `
			SELECT * FROM uploads
			WHERE collection_id = $1
			ORDER BY uploaded_at DESC
			LIMIT $2 
			OFFSET $3 
		`,
	}
}
