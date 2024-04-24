package domain

import (
	"github.com/andrefsilveira1/urban/internal/domain/entity"
)

type UserService struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) *UserService {
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
