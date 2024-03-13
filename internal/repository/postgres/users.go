package repository

import (
	"context"
	"log"
	"tasktracker/internal/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type UsersRepo struct {
	DB *pgx.Conn
}

func NewUsersRepo(conn *pgx.Conn) *UsersRepo {
	return &UsersRepo{
		DB: conn,
	}
}

func (u *UsersRepo) Create(ctx context.Context, user domain.Users) error {
	_, err := u.DB.Exec(ctx, "INSERT INTO users (id, name, password_hash) VALUES ($1, $2, $3)", user.ID, user.Name, user.PasswordHash)
	// TODO : dublicate check
	if err != nil {
		log.Printf("%s", err)
	}
	return err
}

func (u *UsersRepo) GetByCredentials(ctx context.Context, name string) (domain.Users, error) {
	var (
		u_id            string
		u_name          string
		u_password_hash string
	)

	row := u.DB.QueryRow(ctx, "SELECT id, name, password_hash FROM users WHERE name=$1", name)
	err := row.Scan(&u_id, &u_name, &u_password_hash)
	if err != nil {
		return domain.Users{}, err
	}

	u_uuid, _ := uuid.Parse(u_id)
	user := domain.Users{
		ID:           u_uuid,
		Name:         u_name,
		PasswordHash: u_password_hash,
	}

	return user, nil
}

func (u *UsersRepo) GetPasswordHashByUsername(ctx context.Context, name string) (string, error) {
	var password_hash string
	err := u.DB.QueryRow(ctx, "SELECT password_hash FROM users WHERE name=$1", name).Scan(&password_hash)
	if err != nil {
		return "", err
	}
	return password_hash, nil
}

func (u *UsersRepo) GetUserIdByRefreshToken(ctx context.Context, refreshToken string) (uuid.UUID, error) {
	var userID uuid.UUID

	err := u.DB.QueryRow(ctx, "SELECT user_id FROM users_sessions WHERE refresh_token=$1", refreshToken).Scan(&userID)
	if err != nil {
		return uuid.UUID{}, err
	}

	return userID, nil
}

func (u *UsersRepo) SetSession(ctx context.Context, session domain.Session, userID uuid.UUID) error {
	query := `
	INSERT INTO users_sessions (user_id, refresh_token, expires_at)
	VALUES ($1, $2, $3)
	ON CONFLICT (user_id) DO UPDATE
	SET refresh_token = EXCLUDED.refresh_token, expires_at = EXCLUDED.expires_at
	`

	_, err := u.DB.Exec(ctx, query, userID, session.RefreshToken, session.ExpiresAt)
	if err != nil {
		log.Printf("%s", err)
	}
	return err
}
