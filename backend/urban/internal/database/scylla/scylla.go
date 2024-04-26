package scylla

import (
	"log"

	"github.com/andrefsilveira1/urban/internal/config"
	"github.com/gocql/gocql"
)

func Connect(cfg *config.Database) (*gocql.Session, error) {
	cluster := gocql.NewCluster(cfg.Cluster...)
	cluster.Keyspace = cfg.Keyspace
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: cfg.Username,
		Password: cfg.Password,
	}

	session, err := cluster.CreateSession()
	if err != nil {
		log.Printf("scylla connection error: %v", err)
	}

	return session, nil
}
