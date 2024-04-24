package domain

import (
	"github.com/andrefsilveira1/urban/internal/domain/entity"
)

type ImageRepository interface {
	Save(image *entity.Image) error
	Get(id string) (*entity.Image, error)
	List() (*[]entity.Image, error)
	Delete(id string) error
}
