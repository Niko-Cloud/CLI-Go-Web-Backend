package service

import (
	"CliPorto/internal/domain"
	"CliPorto/internal/repository"
	"CliPorto/internal/utils"
)

type ProfileService struct {
	repo *repository.ProfileRepository
}

func NewProfileService(repo *repository.ProfileRepository) *ProfileService {
	return &ProfileService{repo: repo}
}

func (s *ProfileService) GetAll() ([]domain.Profile, error) {
	data, err := s.repo.GetLatest()
	if err != nil {
		return nil, utils.NewBadRequest(err.Error())
	}
	return data, nil
}
