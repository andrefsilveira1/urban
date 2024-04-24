package domain

import (
	"github.com/andrefsilveira1/urban/internal/domain/entity"
)

type UserRepository interface {
	SaveUser(user *entity.User) error
	GetUser(id string) (*entity.User, error)
	DeleteUser(id string) error
}
