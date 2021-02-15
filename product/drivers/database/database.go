package database

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/larien/product-service/product/drivers/config"
	_ "github.com/lib/pq" // postgres side effect
)

// DB is an alias for the database DB that represents the connection with the database
type DB = sqlx.DB

// ErrClosed happens when the database is closed
var ErrClosed = errors.New("sql: database is closed")

// New creates a new connection with the database
func New(dbConfig *config.DB) (*DB, error) {
	info := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)
	var err error
	db, err := sqlx.Open("postgres", info)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
