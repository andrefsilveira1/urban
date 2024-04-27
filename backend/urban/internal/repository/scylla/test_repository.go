package repository

import (
	"fmt"

	"github.com/gocql/gocql"
)

const (
	test = "test scylla"
)

var testQueries = map[string]string{
	test: `SELECT release_version FROM system.local WHERE key = 'local'`,
}

type TestRepository struct {
	DB *gocql.Session
}

func NewTestRepository(db *gocql.Session) *TestRepository {
	return &TestRepository{
		DB: db,
	}
}

func (r *TestRepository) Test() error {
	query := testQueries[test]
	var releaseVersion string
	if err := r.DB.Query(query).Scan(&releaseVersion); err != nil {
		return fmt.Errorf("error testing scylla %w", err)
	}
	fmt.Printf("scylla connected: %s \n", releaseVersion)

	return nil
}
