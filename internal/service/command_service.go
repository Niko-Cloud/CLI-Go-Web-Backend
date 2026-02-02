package service

import (
	"CliPorto/internal/domain"
	"CliPorto/internal/repository"
	"CliPorto/internal/utils"
)

type CommandService struct {
	repo *repository.CommandRepository
}

func NewCommandService(r *repository.CommandRepository) *CommandService {
	return &CommandService{repo: r}
}

func (s *CommandService) GetAll() ([]domain.Command, error) {
	data, err := s.repo.GetAll()
	if err != nil {
		return nil, utils.NewBadRequest(err.Error())
	}
	return data, nil
}

func (s *CommandService) GetByName(name string) (*domain.Command, error) {
	if name == "" {
		return nil, utils.NewBadRequest("command name is required")
	}

	data, err := s.repo.GetByName(name)
	if err != nil {
		return nil, utils.NewBadRequest(err.Error())
	}

	if data == nil {
		return nil, utils.NewNotFound("command not found")
	}

	return data, nil
}
