package scylla

import "github.com/gocql/gocql"

func Connect() {
	cluster := gocql.NewCluster("127.0.0.1") // Replace
	cluster.Keyspace = "urban-keyspace"      // This keyspace does not exist
	cluster.Consistency = gocql.Quorum

	session, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}

	defer session.Close()
}
