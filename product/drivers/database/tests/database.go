package tests

import (
	"log"
	"testing"

	"github.com/larien/product/product/drivers/config"
	"github.com/larien/product/product/drivers/database"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

const (
	password = "secret"
	user     = "postgres"
	dbName   = "postgres"
	host     = "localhost"
)

type Database struct {
	Connection *database.DB
	Port       string
}

// NewDB creates a new container with a stable connection to be used by tests. It will purge
// the creates image and container afterwards.
func NewDB(t *testing.T) *Database {
	t.Helper()
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "9.6",
		Env: []string{
			"POSTGRES_PASSWORD=" + password,
			"POSTGRES_DB=" + dbName,
		},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	var port string
	var db *database.DB
	err = pool.Retry(func() error {
		port = resource.GetPort("5432/tcp")
		var err2 error
		dbConfig := &config.DB{
			User:     user,
			Password: password,
			Host:     host,
			Port:     port,
			Name:     dbName,
		}
		db, err2 = database.New(dbConfig)
		return err2
	})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	t.Cleanup(func() {
		err := pool.Purge(resource)
		if err != nil {
			t.Logf("Could not purge resource: %s", err)
		}
	})
	return &Database{
		Connection: db,
		Port:       port,
	}
}

// DefaultConnection creates a new connection in an existing container
func DefaultConnection(port string) (*database.DB, error) {
	return database.New(&config.DB{
		User:     user,
		Password: password,
		Host:     host,
		Port:     port,
		Name:     dbName,
	})
}
