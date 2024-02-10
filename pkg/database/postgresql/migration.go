package postgresql

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
)

const migrationURL = "file://./db/migrations"

func runDBMigrations(conn *sql.DB, databaseName string) error {
	driver, err := postgres.WithInstance(conn, &postgres.Config{})
	if err != nil {
		panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		migrationURL,
		databaseName,
		driver,
	)
	if err != nil {
		panic(err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		panic(err)
	}

	return nil
}
