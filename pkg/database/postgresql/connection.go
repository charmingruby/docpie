package postgresql

import (
	"database/sql"
	"fmt"

	"github.com/charmingruby/docpie/config"
	_ "github.com/lib/pq"
)

func ConnectDB(cfg *config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=require",
		cfg.Database.DatabaseUser,
		cfg.Database.DatabasePassword,
		cfg.Database.DatabaseHost,
		cfg.Database.DatabaseName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
