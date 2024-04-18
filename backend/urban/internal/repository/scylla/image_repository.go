package repository

import "github.com/andrefsilveira1/urban/internal/domain/models"

type ImageRepository interface {
	Save(image *models.Image) error
	Get(id string) (*models.Image, error)
	List() (*[]models.Image, error)
	// ...
}
