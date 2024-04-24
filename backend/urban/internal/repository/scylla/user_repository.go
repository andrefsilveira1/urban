package scylla

import (
	"fmt"

	"github.com/andrefsilveira1/urban/internal/domain/entity"
	"github.com/gocql/gocql"
)

const (
	createUser = "createUser"
	deleteUser = "deleteUser"
	getUser    = "getUser"
)

var queriesUser = map[string]string{
	createUser: `INSERT INTO users (id, email, name, password) VALUES (?, ?, ?, ?)`,
	deleteUser: `DELETE FROM users WHERE id = ?`,
	getUser:    `SELECT email, name FROM users WHERE id = ? LIMIT 1`,
}

type UserRepository struct {
	DB *gocql.Session
}

func NewUserRepository(db *gocql.Session) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) CreateUser(user entity.User) error {
	query := queriesUser[createUser]
	if err := r.DB.Query(query, user.ID, user.Email, user.Name, user.Password).Exec(); err != nil {
		return fmt.Errorf("error creating user: %w", err)
	}
	return nil
}

func (r *UserRepository) DeleteUser(id gocql.UUID) error {
	query := queriesUser[deleteUser]
	if err := r.DB.Query(query, id).Exec(); err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}
	return nil
}

func (r *UserRepository) GetUser(id gocql.UUID) (*entity.User, error) {
	query := queriesUser[getUser]
	var user entity.User
	if err := r.DB.Query(query, id).Scan(&user.Email, &user.Name); err != nil {
		if err == gocql.ErrNotFound {
			return nil, fmt.Errorf("user not found with ID %s", id)
		}
		return nil, fmt.Errorf("error getting user: %w", err)
	}

	return &user, nil
}
