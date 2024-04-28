package repository

import (
	"fmt"

	"github.com/andrefsilveira1/urban/internal/domain/entity"
	"github.com/gocql/gocql"
)

const (
	createImage = "create image"
	deleteImage = "delete image by id"
	getImage    = "get image by id"
	listImage   = "list images"
)

var queries = map[string]string{
	createImage: `INSERT INTO images (id, name, date, content) VALUES (?, ?, ?, ?)`,
	deleteImage: `DELETE FROM images WHERE id = ?`,
	getImage:    `SELECT id, name, date, content FROM images WHERE id = ? LIMIT 1`,
	listImage:   `SELECT id, name, date, content FROM images`,
}

type ImageRepository struct {
	DB *gocql.Session
}

func NewImageRepository(db *gocql.Session) *ImageRepository {
	return &ImageRepository{
		DB: db,
	}
}

func (r *ImageRepository) Save(image *entity.Image) error {
	query := queries[createImage]
	if err := r.DB.Query(query, image.Id, image.Name, image.Date, image.Content).Exec(); err != nil {
		return fmt.Errorf("error creating image: %w", err)
	}
	return nil
}

func (r *ImageRepository) Delete(id string) error {
	query := queries[deleteImage]
	if err := r.DB.Query(query, id).Exec(); err != nil {
		return fmt.Errorf("error deleting image: %w", err)
	}
	return nil
}

func (r *ImageRepository) Get(id string) (*entity.Image, error) {
	query := queries[getImage]
	var img entity.Image
	if err := r.DB.Query(query, id).Scan(&img.Id, &img.Name, &img.Date, &img.Content); err != nil {
		if err == gocql.ErrNotFound {
			return nil, fmt.Errorf("image not found with ID %s", id)
		}
		return nil, fmt.Errorf("error getting image: %w", err)
	}
	return &img, nil
}

func (r *ImageRepository) List() ([]*entity.Image, error) {
	query := queries[listImage]
	var images []*entity.Image
	iter := r.DB.Query(query).Iter()
	defer iter.Close()

	for {
		var img *entity.Image
		if !iter.Scan(&img.Id, &img.Name, &img.Date, &img.Content) {
			break
		}
		images = append(images, img)
	}

	if err := iter.Close(); err != nil {
		return nil, fmt.Errorf("error listing images: %w", err)
	}

	return images, nil
}
