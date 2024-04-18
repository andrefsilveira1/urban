package domain

import (
	"time"

	"github.com/andrefsilveira1/urban/internal/domain/models"
	repository "github.com/andrefsilveira1/urban/internal/repository/scylla"
)

type ImageService struct {
	imageRepository repository.ImageRepository
}

func (s *ImageService) Register(name string, date time.Time, content []byte) error {
	image := &models.Image{
		Name:    name,
		Date:    date,
		Content: content,
	}
	err := s.imageRepository.Save(image)
	if err != nil {
		return err
	}

	return nil
}

func (s *ImageService) Get(id string) (*models.Image, error) {
	image, err := s.imageRepository.Get(id)
	if err != nil {
		return nil, err
	}

	return image, nil
}
