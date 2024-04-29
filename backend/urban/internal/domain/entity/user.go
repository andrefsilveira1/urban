package entity

import "github.com/gocql/gocql"

type User struct {
	Id       gocql.UUID `json:"id"`
	Name     string     `json:"name"`
	Email    string     `json:"email"`
	Password string     `json:"password"` // this should be changed latter
}
