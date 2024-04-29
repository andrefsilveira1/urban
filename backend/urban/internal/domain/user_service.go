package domain

import (
	"errors"

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

func (s *UserService) Register(user *entity.User) error {
	if user == nil {
		return errors.New("user can't be null")
	}
	return s.userRepository.SaveUser(user)
}

func (s *UserService) Get(id string) (*entity.User, error) {
	if id == "" {
		return nil, errors.New("id is null")
	}
	return s.userRepository.GetUser(id)

}

func (s *UserService) Delete(id string) error {
	if id == "" {
		return errors.New("id is null")
	}

	return s.userRepository.DeleteUser(id)
}
