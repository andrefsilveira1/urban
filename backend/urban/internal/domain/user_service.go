package domain

type UserService struct {
	userRepository UserRepository
}

func (s *UserService) Register(name, email, password string) (*User, error) {
	user := &User{
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
