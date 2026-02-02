package service

import (
	"CliPorto/internal/domain"
	"CliPorto/internal/repository"
)

type ProfileService struct {
	repo *repository.ProfileRepository
}

func NewProfileService(repo *repository.ProfileRepository) *ProfileService {
	return &ProfileService{repo: repo}
}

func (s *ProfileService) GetAll() ([]domain.Profile, error) {
	return s.repo.GetLatest()
}
