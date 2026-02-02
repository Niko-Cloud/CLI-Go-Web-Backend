package service

import (
	"CliPorto/internal/domain"
	"CliPorto/internal/repository"
)

type ContactService struct {
	repo *repository.ContactRepository
}

func NewContactService(repo *repository.ContactRepository) *ContactService {
	return &ContactService{repo: repo}
}

func (s *ContactService) GetContactInfo() ([]domain.Contact, error) {
	return s.repo.GetContactInfo()
}
