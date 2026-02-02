package service

import (
	"CliPorto/internal/domain"
	"CliPorto/internal/repository"
)

type SkillService struct {
	repo *repository.SkillRepository
}

func NewSkillService(repo *repository.SkillRepository) *SkillService {
	return &SkillService{repo: repo}
}

func (s *SkillService) GetAll() ([]domain.Skill, error) {
	return s.repo.GetAll()
}
