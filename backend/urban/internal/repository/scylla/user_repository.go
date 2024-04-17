package repository

import (
	"github.com/andrefsilveira1/urban/internal/domain/models"
)

type UserRepository interface {
	Save(user *models.User) error
	Get(id string) (models.User, error)
	// more methods
}

type ScyllaRepository struct {
	// Scylla init
}

func (r *ScyllaRepository) Save(User *models.User) error {
	// Save logic

	return nil
}

func (r *ScyllaRepository) Get(id string) (*models.User, error) {
	// Get logic

	return nil, nil
}
