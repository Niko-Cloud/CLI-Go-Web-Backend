package repository

import (
	"CliPorto/internal/domain"
	"database/sql"
)

type EducationRepository struct {
	db *sql.DB
}

func NewEducationRepository(db *sql.DB) *EducationRepository {
	return &EducationRepository{db: db}
}

func (r *EducationRepository) GetAll() ([]domain.Education, error) {
	const query = `
		SELECT
			id,
			institution,
			degree,
			major,
			start_date,
			end_date,
			description
		FROM education
		ORDER BY start_date DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.Education

	for rows.Next() {
		var e domain.Education
		if err := rows.Scan(
			&e.ID,
			&e.Institution,
			&e.Degree,
			&e.Major,
			&e.StartDate,
			&e.EndDate,
			&e.Description,
		); err != nil {
			return nil, err
		}
		result = append(result, e)
	}

	return result, rows.Err()
}
