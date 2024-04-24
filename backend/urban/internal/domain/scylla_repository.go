package domain

import "github.com/gocql/gocql"

type ScyllaRepository struct {
	session *gocql.Session
}
