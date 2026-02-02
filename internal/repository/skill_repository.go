package repository

import (
	"CliPorto/internal/domain"
	"database/sql"
	"github.com/lib/pq"
)

type SkillRepository struct {
	db *sql.DB
}

func NewSkillRepository(db *sql.DB) *SkillRepository {
	return &SkillRepository{db: db}
}

func (r *SkillRepository) GetAll() ([]domain.Skill, error) {
	const query = `
		SELECT
			id,
			name,
			level,
			category
		FROM skill
		ORDER BY id ASC
		`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []domain.Skill

	for rows.Next() {
		var s domain.Skill
		if err := rows.Scan(
			&s.ID,
			&s.Name,
			&s.Level,
			pq.Array(&s.Category),
		); err != nil {
			return nil, err
		}

		results = append(results, s)
	}

	return results, rows.Err()
}
