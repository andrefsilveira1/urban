package repository

import (
	"fmt"

	"github.com/andrefsilveira1/urban/internal/domain/models"
)

type UserRepository interface {
	SaveUser(user *models.User) error
	GetUser(id string) (*models.User, error)
	ListUsers() (*[]models.User, error)
	// more methods
}

func (r *ScyllaRepository) SaveUser(User *models.User) error {
	query := "INSERT INTO users (id, email, name, password) VALUES (?, ?, ?, ?)"
	if err := r.session.Query(query, User.ID, User.Email, User.Name, User.Password).Exec(); err != nil {
		return fmt.Errorf("error: saving user has failed: %v", err)
	}

	return nil
}

func (r *ScyllaRepository) GetUser(id string) (*models.User, error) {
	var user models.User

	query := "SELECT id, email, name FROM users WHERE id = ? LIMIT 1"
	if err := r.session.Query(query, id).Scan(&user.ID, &user.Email, &user.Name); err != nil {
		return nil, fmt.Errorf("error: saving image has failed: %v", err)
	}
	// Create authorization later
	return &user, nil

}
