package repository

import (
	"fmt"

	"github.com/andrefsilveira1/urban/internal/domain/models"
	"github.com/gocql/gocql"
)

type ImageRepository interface {
	Save(image *models.Image) error
	Get(id string) (*models.Image, error)
	List() (*[]models.Image, error)
	Delete(id string) error
}

func NewScyllaImageRepository(session *gocql.Session) *ScyllaRepository {
	return &ScyllaRepository{
		session: session,
	}
}

func (r *ScyllaRepository) Save(image *models.Image) error {
	query := "INSERT INTO images (id, name, date, content) VALUES (?, ?, ?, ?)"
	if err := r.session.Query(query, image.Id, image.Name, image.Date, image.Content).Exec(); err != nil {
		return fmt.Errorf("Error: Saving image has failed: $v", err)
	}
	return nil
}

func (r *ScyllaRepository) Get(id string) (*models.Image, error) {
	var image models.Image

	query := "SELECT id, name, date, content FROM images WHERE id = ? LIMIT 1"
	if err := r.session.Query(query, id).Scan(&image.Id, &image.Name, &image.Date, &image.Content); err != nil {
		return nil, fmt.Errorf("Error: Saving image has failed: $v", err)
	}

	return &image, nil
}

func (r *ScyllaRepository) List() ([]models.Image, error) {
	var images []models.Image
	query := "SELECT id, name, date, content FROM images"
	iter := r.session.Query(query).Iter()
	defer iter.Close()

	var id, name string
	var date gocql.UUID
	var content []byte

	for iter.Scan(&id, &name, &date, &content) {
		images = append(images, models.Image{
			Id:      id,
			Name:    name,
			Date:    date.Time(),
			Content: content,
		})
	}

	if err := iter.Close(); err != nil {
		return nil, fmt.Errorf("Error while listing images: %v", err)
	}

	return images, nil

}

func (r *ScyllaRepository) Delete(id string) error {
	query := "DELETE FROM images WHERE id = ?"
	if err := r.session.Query(query, id).Exec(); err != nil {
		return fmt.Errorf("error deleting image: %v", err)
	}
	return nil
}
