package db

import (
	"context"
	"database/sql"

	"github.com/dayachettri/hotel-reservation/types"
)

type UserStore interface {
	GetUserByID(context.Context, string) (*types.User, error)
	GetUsers(context.Context) ([]*types.User, error)
	CreateUser(context.Context, *types.User) (*types.User, error)
}

type PostgresUserStore struct {
	db *sql.DB
}

func NewPostgresUserStore(db *sql.DB) *PostgresUserStore {
	return &PostgresUserStore{
		db: db,
	}
}

func (s *PostgresUserStore) CreateUser(ctx context.Context, u *types.User) (*types.User, error) {
	query := `INSERT INTO users (first_name, last_name, email, encrypted_password)
              VALUES ($1, $2, $3, $4)
              RETURNING id`

	stmt, err := s.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	if err := stmt.QueryRow(u.FirstName, u.LastName, u.Email, u.EncryptedPassword).Scan(&u.ID); err != nil {
		return nil, err
	}

	return u, nil
}

func (s *PostgresUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	user := types.User{}

	query := `SELECT * FROM users WHERE id = $1`

	stmt, err := s.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(query, id)

	if err := row.Scan(&user.ID, &user.FirstName, &user.LastName); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *PostgresUserStore) GetUsers(ctx context.Context) ([]*types.User, error) {
	users := []*types.User{}

	query := `SELECT id, first_name, last_name, email FROM users`

	stmt, err := s.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return users, nil
	}

	for rows.Next() {
		user := &types.User{}

		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, err
}
