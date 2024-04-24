package domain

import (
	"errors"

	"github.com/andrefsilveira1/urban/internal/domain/entity"
)

type ImageService struct {
	imageRepository ImageRepository
}

func NewImageService(repo ImageRepository) *ImageService {
	return &ImageService{
		imageRepository: repo,
	}
}

func (s *ImageService) Register(image *entity.Image) error {

	if image == nil {
		return errors.New("image can't be null")
	}
	return s.imageRepository.Save(image)
}

func (s *ImageService) Get(id string) (*entity.Image, error) {
	if id == "" {
		return nil, errors.New("id can't be null")
	}

	return s.imageRepository.Get(id)

}

func (s *ImageService) List() (*[]entity.Image, error) {
	return s.imageRepository.List()

}

func (s *ImageService) Delete(id string) error {
	if id == "" {
		return errors.New("invalid id")
	}
	return s.imageRepository.Delete(id)

}
