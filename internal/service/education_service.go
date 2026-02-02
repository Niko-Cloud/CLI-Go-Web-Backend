package service

import (
	"CliPorto/internal/domain"
	"CliPorto/internal/repository"
	"CliPorto/internal/utils"
)

type EducationService struct {
	repo *repository.EducationRepository
}

func NewEducationService(r *repository.EducationRepository) *EducationService {
	return &EducationService{repo: r}
}

func (s *EducationService) GetAll() ([]domain.Education, error) {
	data, err := s.repo.GetAll()
	if err != nil {
		return nil, utils.NewBadRequest(err.Error())
	}
	return data, nil
}
