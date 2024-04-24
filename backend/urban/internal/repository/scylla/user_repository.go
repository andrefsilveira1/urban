package scylla

import (
	"github.com/andrefsilveira1/urban/internal/domain/entity"
	"github.com/gocql/gocql"
)

const (
	createUser = "create image"
	deleteUser = "delete image by id"
	getUser    = "get image by id"
)

var queriesUser = map[string]string{
	createUser: `INSERT INTO users (id, email, name, password) VALUES (?, ?, ?, ?)`,
	deleteUser: `DELETE FROM user WHERE id = ?`,
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
	return r.DB.Query(queriesUser[createUser], user.ID, user.Email, user.Name, user.Password).Exec()
}

func (r *UserRepository) DeleteUser(id gocql.UUID) error {
	return r.DB.Query(queriesUser[deleteUser], id).Exec()
}

func (r *UserRepository) GetUser(id gocql.UUID) (*entity.User, error) {
	var user entity.User
	if err := r.DB.Query(queriesUser[getUser], id).Scan(&user.Email, &user.Name); err != nil {
		return nil, err
	}

	return &user, nil
}
