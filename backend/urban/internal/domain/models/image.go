package models

import "time"

type Image struct {
	Name    string    `json:"name"`    // Name of the image
	Date    time.Time `json:"date"`    // Date when the image was created
	Content []byte    `json:"content"` // Binary data of the image
}
