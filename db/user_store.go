package db

import (
	"database/sql"

	"github.com/dayachettri/hotel-reservation/types"
	"github.com/labstack/echo/v4"
)

type UserStore interface {
	GetUserByID(string) (*types.User, error)
	GetUsers() ([]*types.User, error)
}

type PostgresUserStore struct {
	db *sql.DB
}

func NewPostgresUserStore(db *sql.DB) *PostgresUserStore {
	return &PostgresUserStore{
		db: db,
	}
}

func (s *PostgresUserStore) GetUserByID(id string) (*types.User, error) {
	user := types.User{}

	query := `SELECT * FROM users WHERE id = $1`

	row := s.db.QueryRow(query, id)

	if err := row.Scan(&user.ID, &user.FirstName, &user.LastName); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *PostgresUserStore) GetUsers() ([]*types.User, error) {
	users := []*types.User{}

	query := `SELECT id, first_name, last_name FROM users`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, echo.ErrInternalServerError
	}
	defer rows.Close()

	if !rows.Next() {
		return users, nil
	}

	for rows.Next() {
		user := &types.User{}

		err := rows.Scan(&user.FirstName, &user.LastName, &user.ID)
		if err != nil {
			return nil, echo.ErrInternalServerError
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, echo.ErrInternalServerError
	}

	return users, err
}
