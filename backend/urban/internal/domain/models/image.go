package models

import "time"

type Image struct {
	Id      string    `json:"id"`
	Name    string    `json:"name"`
	Date    time.Time `json:"date"`
	Content []byte    `json:"content"`
}
