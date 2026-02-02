package repository

import (
	"CliPorto/internal/domain"
	"database/sql"
	"errors"
	"github.com/lib/pq"
)

type ShowcaseRepository struct {
	db *sql.DB
}

func NewShowcaseRepository(db *sql.DB) *ShowcaseRepository {
	return &ShowcaseRepository{db: db}
}

func (r *ShowcaseRepository) GetAll() ([]domain.Showcase, error) {
	const query = `
		SELECT
			id,
			title,
			summary,
			stack,
			repo_url,
			live_url
		FROM showcase
		ORDER BY id ASC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []domain.Showcase

	for rows.Next() {
		var s domain.Showcase
		if err := rows.Scan(
			&s.ID,
			&s.Title,
			&s.Summary,
			pq.Array(&s.Stack),
			&s.RepoURL,
			&s.LiveURL,
		); err != nil {
			return nil, err
		}
		results = append(results, s)
	}

	return results, rows.Err()
}

func (r *ShowcaseRepository) GetByID(id int) (*domain.Showcase, error) {
	const query = `
		SELECT
			id,
			title,
			summary,
			stack,
			repo_url,
			live_url
		FROM showcase
		WHERE id = $1
	`

	var s domain.Showcase

	err := r.db.QueryRow(query, id).Scan(
		&s.ID,
		&s.Title,
		&s.Summary,
		pq.Array(&s.Stack),
		&s.RepoURL,
		&s.LiveURL,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &s, nil
}
