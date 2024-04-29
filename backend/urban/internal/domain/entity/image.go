package entity

import (
	"time"

	"github.com/gocql/gocql"
)

type Image struct {
	Id      gocql.UUID `json:"id"`
	Name    string     `json:"name"`
	Date    time.Time  `json:"date"`
	Content []byte     `json:"content"`
}
