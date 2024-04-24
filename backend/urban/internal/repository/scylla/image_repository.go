package scylla

import (
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

func (r *ImageRepository) CreateImage(id gocql.UUID, name, date gocql.UUID, content []byte) error {
	return r.DB.Query(queries[createImage], id, name, date, content).Exec()
}

func (r *ImageRepository) DeleteImage(id gocql.UUID) error {
	return r.DB.Query(queries[deleteImage], id).Exec()
}

func (r *ImageRepository) GetImageById(id gocql.UUID) (*entity.Image, error) {
	var img entity.Image
	if err := r.DB.Query(queries[getImage], id).Scan(&img.Id, &img.Name, &img.Date, &img.Content); err != nil {
		return nil, err
	}

	return &img, nil
}

func (r *ImageRepository) ListImages() ([]entity.Image, error) {
	var imgs []entity.Image
	iter := r.DB.Query(queries[listImage]).Iter()
	defer iter.Close()

	for {
		var img entity.Image
		if !iter.Scan(&img.Id, &img.Name, &img.Date, &img.Content) {
			break
		}
		imgs = append(imgs, img)
	}

	return imgs, nil

}
