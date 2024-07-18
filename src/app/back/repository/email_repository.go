package repository

import (
	"checkit/models"
	"database/sql"
)

type EmailRepository struct {
	db *sql.DB
}

func NewEmailRepository(db *sql.DB) *EmailRepository {
	return &EmailRepository{db: db}
}

func (r *EmailRepository) GetEmails() ([]models.Email, error) {
	return nil, nil
	// Implement database query to get all emails
}

func (r *EmailRepository) GetEmail(id string) (*models.Email, error) {
	return nil, nil
	// Implement database query to get a single email
}

// Implement other CRUD methods...
