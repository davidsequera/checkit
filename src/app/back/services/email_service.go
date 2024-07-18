package services

import (
	"checkit/models"
	"checkit/repository"
)

type EmailService struct {
	repo *repository.EmailRepository
}

func NewEmailService(repo *repository.EmailRepository) *EmailService {
	return &EmailService{repo: repo}
}

func (s *EmailService) GetEmails() ([]models.Email, error) {
	return s.repo.GetEmails()
}

func (s *EmailService) GetEmail(id string) (*models.Email, error) {
	return s.repo.GetEmail(id)
}

// Implement other business logic methods...
