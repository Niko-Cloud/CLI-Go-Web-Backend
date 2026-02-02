package repository

import (
	"CliPorto/internal/domain"
	"CliPorto/internal/utils"
	"database/sql"
	"errors"
)

type ProfileRepository struct {
	db *sql.DB
}

func NewProfileRepository(db *sql.DB) *ProfileRepository {
	return &ProfileRepository{db: db}
}

func (r *ProfileRepository) GetLatest() ([]domain.Profile, error) {
	const query = `
		SELECT
			id,
			name,
			role,
			description,
			location
		FROM profile
		WHERE created_at = (
			SELECT MAX(updated_at) FROM profile
		)
		ORDER BY id ASC LIMIT 1
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []domain.Profile

	for rows.Next() {
		var p domain.Profile
		if err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Role,
			&p.Description,
			&p.Location,
		); err != nil {
			return nil, err
		}
		results = append(results, p)
	}

	if err := rows.Err(); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, utils.ErrNotFound
		}
		return nil, err
	}

	return results, nil
}
