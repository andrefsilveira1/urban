package domain

import (
	"time"

	"github.com/andrefsilveira1/urban/internal/domain/entity"
	"github.com/gocql/gocql"
)

type Events struct {
	Id     gocql.UUID     `json:"id"`
	Name   string         `json:"name"`
	Date   time.Time      `json:"date"`
	Images []entity.Image `json:"images"`
}
