package service

import (
	"CliPorto/internal/domain"
	"CliPorto/internal/repository"
	"CliPorto/internal/utils"
)

type SkillService struct {
	repo *repository.SkillRepository
}

func NewSkillService(repo *repository.SkillRepository) *SkillService {
	return &SkillService{repo: repo}
}

func (s *SkillService) GetAll() ([]domain.Skill, error) {
	data, err := s.repo.GetAll()
	if err != nil {
		return nil, utils.NewBadRequest(err.Error())
	}
	return data, nil
}
