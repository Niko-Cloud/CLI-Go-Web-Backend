package database

import (
	"CliPorto/internal/config"
	"database/sql"

	_ "github.com/lib/pq"
)

func Connect(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.Database.URI)
	if err != nil {
		return nil, err
	}
	return db, db.Ping()
}
