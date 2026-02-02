package service

import (
	"CliPorto/internal/domain"
	"CliPorto/internal/repository"
	"CliPorto/internal/utils"
)

type ContactService struct {
	repo *repository.ContactRepository
}

func NewContactService(repo *repository.ContactRepository) *ContactService {
	return &ContactService{repo: repo}
}

func (s *ContactService) GetContactInfo() ([]domain.Contact, error) {
	data, err := s.repo.GetContactInfo()
	if err != nil {
		return nil, utils.NewBadRequest(err.Error())
	}
	return data, nil
}
