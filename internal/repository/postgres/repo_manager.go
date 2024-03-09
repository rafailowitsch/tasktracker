package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func CreateTables(ctx context.Context, conn *pgx.Conn) error {
	var err error

	_, err = conn.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY,
			name VARCHAR(255) UNIQUE,
			password_hash VARCHAR(255)
		)
	`)
	if err != nil {
		return err
	}

	_, err = conn.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS users_session (
			user_id UUID PRIMARY KEY,
			refresh_token VARCHAR(255) NOT NULL,
			expires_at TIMESTAMP NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id)
		)
	`)
	if err != nil {
		return err
	}

	return nil
}
