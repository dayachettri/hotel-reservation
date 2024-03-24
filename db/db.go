package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type Database interface {
	Connect(string) error
}

type PostgresStore struct {
	DB *sql.DB
}

func NewPostgresStore() *PostgresStore {
	return &PostgresStore{}
}

func (d *PostgresStore) Connect(url string) error {
	databaseURL := os.Getenv(url)
	if databaseURL == "" {
		return errors.New("missing DATABASE_URL environment variable")
	}

	database, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return err
	}

	if err = database.Ping(); err != nil {
		return err
	}

	fmt.Println("connection established to database")
	d.DB = database
	return nil
}
