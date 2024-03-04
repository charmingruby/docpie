package postgresql

import (
	"fmt"

	"github.com/charmingruby/upl/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func LoadDatabase(cfg *config.Config) (*sqlx.DB, error) {
	connStr := fmt.Sprintf(
		"postgresql://%s:%s@%s/%s?sslmode=%s",
		cfg.Database.DatabaseUser,
		cfg.Database.DatabasePassword,
		cfg.Database.DatabaseHost,
		cfg.Database.DatabaseName,
		cfg.Database.DatabaseSSL,
	)

	db, err := sqlx.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
