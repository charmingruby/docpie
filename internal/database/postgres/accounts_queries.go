package postgres

const (
	createAccount      = "create account"
	findAccountByEmail = "find by email"
	findAccountById    = "find by id"
	saveAccount        = "save account"
)

func accountQueries() map[string]string {
	return map[string]string{
		createAccount:      `INSERT INTO accounts (id, name, last_name, email, role, password, avatar_url, upload_quantity, collections_member_quantity, collections_created_quantity) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING *`,
		findAccountByEmail: `SELECT * FROM accounts WHERE email = $1`,
		findAccountById:    `SELECT * FROM accounts WHERE id = $1`,
		saveAccount:        `UPDATE accounts SET name = $1, last_name = $2, email = $3, role = $4, avatar_url = $5, updated_at = $6, deleted_by = $7, deleted_at = $8, upload_quantity = $9, collections_member_quantity= $10, collections_created_quantity = $11 where id = $12`,
	}
}
