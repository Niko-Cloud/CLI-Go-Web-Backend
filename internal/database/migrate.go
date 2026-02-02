package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("migration driver error: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migration",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatalf("migration init error: %v", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("migration failed: %v", err)
	}

	fmt.Println("Database migrated successfully")
}
