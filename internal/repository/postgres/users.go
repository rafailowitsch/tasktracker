package repository

import (
	"context"
	"log"
	"tasktracker/internal/domain"

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
	_, err := u.DB.Exec(context.Background(), "INSERT INTO users (name, password_hash) VALUES ($1, $2) RETURNING id", user.Name, user.Password)
	// TODO : dublicate check
	if err != nil {
		log.Printf("%s", err)
	}
	return err
}

func (u *UsersRepo) GetByCredentials(ctx context.Context, username, password string) (domain.Users, error) {
	user := domain.Users{}
	return user, nil
}
