package service

import (
	"CliPorto/internal/domain"
	"CliPorto/internal/repository"
	"CliPorto/internal/utils"
)

type WorkExperienceService struct {
	repo *repository.WorkExperienceRepository
}

func NewWorkExperienceService(r *repository.WorkExperienceRepository) *WorkExperienceService {
	return &WorkExperienceService{repo: r}
}

func (s *WorkExperienceService) GetAll() ([]domain.WorkExperience, error) {
	data, err := s.repo.GetAll()
	if err != nil {
		return nil, utils.NewBadRequest(err.Error())
	}
	return data, nil
}

func (s *WorkExperienceService) GetByID(id int) (*domain.WorkExperience, error) {
	if id <= 0 {
		return nil, utils.NewBadRequest("invalid ID")
	}

	data, err := s.repo.GetByID(id)
	if err != nil {
		return nil, utils.NewBadRequest(err.Error())
	}

	if data == nil {
		return nil, utils.NewNotFound("work experience not found")
	}

	return data, nil
}
