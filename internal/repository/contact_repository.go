package repository

import (
	"CliPorto/internal/domain"
	"CliPorto/internal/utils"
	"database/sql"
	"errors"
)

type ContactRepository struct {
	db *sql.DB
}

func NewContactRepository(db *sql.DB) *ContactRepository {
	return &ContactRepository{db: db}
}

func (r *ContactRepository) GetContactInfo() ([]domain.Contact, error) {
	const query = `
		SELECT id, 
		       "type", 
		       value, 
		FROM contact
		ORDER BY id
		`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var results []domain.Contact

	for rows.Next() {
		var c domain.Contact
		if err := rows.Scan(
			&c.ID,
			&c.Type,
			&c.Value,
		); err != nil {
			return nil, err
		}

		results = append(results, c)
	}

	if err := rows.Err(); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, utils.ErrNotFound
		}
		return nil, err
	}

	return results, rows.Err()
}
