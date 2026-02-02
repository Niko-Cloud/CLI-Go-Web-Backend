package repository

import (
	"CliPorto/internal/domain"
	"database/sql"
	"errors"
	"github.com/lib/pq"
)

type WorkExperienceRepository struct {
	db *sql.DB
}

func NewWorkExperienceRepository(db *sql.DB) *WorkExperienceRepository {
	return &WorkExperienceRepository{db: db}
}

func (r *WorkExperienceRepository) GetAll() ([]domain.WorkExperience, error) {
	const query = `
		SELECT
			id,
			company,
			position,
			start_date,
			end_date,
			description,
			tech_stack
		FROM work_experience
		ORDER BY start_date DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.WorkExperience

	for rows.Next() {
		var w domain.WorkExperience
		if err := rows.Scan(
			&w.ID,
			&w.Company,
			&w.Position,
			&w.StartDate,
			&w.EndDate,
			pq.Array(&w.Description),
			pq.Array(&w.TechStack),
		); err != nil {
			return nil, err
		}
		result = append(result, w)
	}

	return result, rows.Err()
}

func (r *WorkExperienceRepository) GetByID(id int) (*domain.WorkExperience, error) {
	const query = `
		SELECT
			id,
			company,
			position,
			start_date,
			end_date,
			description,
			tech_stack
		FROM work_experience
		WHERE id = $1
	`

	var w domain.WorkExperience

	err := r.db.QueryRow(query, id).Scan(
		&w.ID,
		&w.Company,
		&w.Position,
		&w.StartDate,
		w.EndDate,
		pq.Array(w.Description),
		pq.Array(w.TechStack),
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &w, nil
}
