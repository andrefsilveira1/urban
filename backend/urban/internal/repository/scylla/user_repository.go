package repository

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}

type UserRepository interface {
	Save(user *User) error
	Get(id string) (*User, error)
	// more methods
}

type ScyllaRepository struct {
	// Scylla init
}

func (r *ScyllaRepository) Save(User *User) error {
	// Save logic

	return nil
}

func (r *ScyllaRepository) Get(id string) (*User, error) {
	// Get logic

	return nil, nil
}
