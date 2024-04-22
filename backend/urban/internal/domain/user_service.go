package domain

import (
	"github.com/andrefsilveira1/urban/internal/domain/models"
	repository "github.com/andrefsilveira1/urban/internal/repository/scylla"
)

type UserService struct {
	userRepository repository.UserRepository
}

func (s *UserService) Register(name string, email string, password string) error {
	user := &models.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	err := s.userRepository.SaveUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) Get(id string) (*models.User, error) {
	user, err := s.userRepository.GetUser(id)
	if err != nil {
		return nil, err
	}

	return user, nil

}

func (s *UserService) List() (*[]models.User, error) {
	users, err := s.userRepository.ListUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}
