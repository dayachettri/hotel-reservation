package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type DB interface {
	Connect(string) (*sql.DB, error)
}

type PostgresStore struct{}

func NewPostgresStore() *PostgresStore {
	return &PostgresStore{}
}

func (d *PostgresStore) Connect(url string) (*sql.DB, error) {
	databaseURL := os.Getenv(url)
	if databaseURL == "" {
		return nil, errors.New("missing DATABASE_URL environment variable")
	}

	database, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err = database.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("connection established to database")

	return database, nil
}
