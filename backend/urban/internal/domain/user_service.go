package domain

import (
	"github.com/andrefsilveira1/urban/internal/domain/models"
	repository "github.com/andrefsilveira1/urban/internal/repository/scylla"
)

type UserService struct {
	userRepository repository.UserRepository
}

func (s *UserService) Register(name, email, password string) (*models.User, error) {
	user := &models.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	err := s.userRepository.Save(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
