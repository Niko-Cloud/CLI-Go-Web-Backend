package repository

import (
	"CliPorto/internal/domain"
	"database/sql"
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
		       value
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

	return results, rows.Err()
}
