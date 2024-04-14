package repository

type UserRepository interface {
	Save(user *User) error
	Get(id string) (*User, error)
	// more methods
}
