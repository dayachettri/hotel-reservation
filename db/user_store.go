package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/dayachettri/hotel-reservation/types"
)

type UserStore interface {
	GetUserByID(context.Context, string) (*types.User, error)
	GetUsers(context.Context) ([]*types.User, error)
	CreateUser(context.Context, *types.User) (*types.User, error)
	DeleteUser(context.Context, string) error
	UpdateUser(context.Context, *types.UpdateUserParams, string) error
}

type PostgresUserStore struct {
	db *sql.DB
}

func NewPostgresUserStore(db *sql.DB) *PostgresUserStore {
	return &PostgresUserStore{
		db: db,
	}
}

func (s *PostgresUserStore) UpdateUser(ctx context.Context, params *types.UpdateUserParams, id string) error {
	query := `UPDATE users SET first_name = $1, last_name = $2 WHERE id = $3`

	stmt, err := s.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, params.FirstName, params.LastName, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

func (s *PostgresUserStore) DeleteUser(ctx context.Context, id string) error {
	query := `DELETE FROM users WHERE id = $1`

	stmt, err := s.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user with id %s not found", id)
	}

	return nil
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

	row := stmt.QueryRow(id)
	if err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.EncryptedPassword); err != nil {
		fmt.Println(err)
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
