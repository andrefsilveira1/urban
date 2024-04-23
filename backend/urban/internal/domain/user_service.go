package domain

import (
	"github.com/andrefsilveira1/urban/internal/domain/entity"
	repository "github.com/andrefsilveira1/urban/internal/repository/scylla"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) Register(name string, email string, password string) error {
	user := &entity.User{
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

func (s *UserService) Get(id string) (*entity.User, error) {
	user, err := s.userRepository.GetUser(id)
	if err != nil {
		return nil, err
	}

	return user, nil

}
