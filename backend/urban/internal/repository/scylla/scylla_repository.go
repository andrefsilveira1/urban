package repository

import "github.com/gocql/gocql"

type ScyllaRepository struct {
	session *gocql.Session
}
