package domain

import (
	"fmt"

	"github.com/andrefsilveira1/urban/internal/domain/entity"
	"github.com/gocql/gocql"
)

type ImageRepository interface {
	Save(image *entity.Image) error
	Get(id string) (*entity.Image, error)
	List() (*[]entity.Image, error)
	Delete(id string) error
}

func NewScyllaImageRepository(session *gocql.Session) *ScyllaRepository {
	return &ScyllaRepository{
		session: session,
	}
}

func (r *ScyllaRepository) Save(image *entity.Image) error {
	query := "INSERT INTO images (id, name, date, content) VALUES (?, ?, ?, ?)"
	if err := r.session.Query(query, image.Id, image.Name, image.Date, image.Content).Exec(); err != nil {
		return fmt.Errorf("error: saving image has failed: %v", err)
	}
	return nil
}

func (r *ScyllaRepository) Get(id string) (*entity.Image, error) {
	var image entity.Image

	query := "SELECT id, name, date, content FROM images WHERE id = ? LIMIT 1"
	if err := r.session.Query(query, id).Scan(&image.Id, &image.Name, &image.Date, &image.Content); err != nil {
		return nil, fmt.Errorf("error: saving image has failed: %v", err)
	}

	return &image, nil
}

func (r *ScyllaRepository) List() ([]entity.Image, error) {
	var images []entity.Image
	query := "SELECT id, name, date, content FROM images"
	iter := r.session.Query(query).Iter()
	defer iter.Close()

	var id, name string
	var date gocql.UUID
	var content []byte

	for iter.Scan(&id, &name, &date, &content) {
		images = append(images, entity.Image{
			Id:      id,
			Name:    name,
			Date:    date.Time(),
			Content: content,
		})
	}

	if err := iter.Close(); err != nil {
		return nil, fmt.Errorf("error while listing images: %v", err)
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
