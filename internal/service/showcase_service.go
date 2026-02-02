package service

import (
	"CliPorto/internal/domain"
	"CliPorto/internal/repository"
)

type ShowcaseService struct {
	repo *repository.ShowcaseRepository
}

func NewShowcaseService(repo *repository.ShowcaseRepository) *ShowcaseService {
	return &ShowcaseService{repo: repo}
}

func (s *ShowcaseService) GetAll() ([]domain.Showcase, error) {
	return s.repo.GetAll()
}

func (s *ShowcaseService) GetByID(id int) (*domain.Showcase, error) {
	return s.repo.GetByID(id)
}
