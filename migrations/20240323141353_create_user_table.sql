-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users(
 id SERIAL PRIMARY KEY,
 first_name VARCHAR(255) NOT NULL,
 last_name VARCHAR(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd