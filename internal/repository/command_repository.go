package repository

import (
	"CliPorto/internal/domain"
	"database/sql"
	"errors"
)

type CommandRepository struct {
	db *sql.DB
}

func NewCommandRepository(db *sql.DB) *CommandRepository {
	return &CommandRepository{db: db}
}

func (r *CommandRepository) GetAll() ([]domain.Command, error) {
	const q = `
		SELECT 
		    id, 
		    name, 
		    description, 
		    usage, 
		    help_text
		FROM commands
		WHERE enabled = TRUE
		ORDER BY name
	`

	rows, err := r.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cmds []domain.Command
	for rows.Next() {
		var c domain.Command
		if err := rows.Scan(
			&c.ID,
			&c.Name,
			&c.Description,
			&c.Usage,
			&c.HelpText,
		); err != nil {
			return nil, err
		}
		cmds = append(cmds, c)
	}

	return cmds, rows.Err()
}

func (r *CommandRepository) GetByName(name string) (*domain.Command, error) {
	const q = `
		SELECT 
		    id, 
		    name, 
		    description, 
		    usage, 
		    help_text
		FROM commands
		WHERE name = $1 AND enabled = TRUE
	`

	var c domain.Command
	err := r.db.QueryRow(q, name).Scan(
		&c.ID,
		&c.Name,
		&c.Description,
		&c.Usage,
		&c.HelpText,
	)

	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &c, nil
}
