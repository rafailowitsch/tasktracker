package repository

import (
	"context"
	"tasktracker/internal/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type Users interface {
	Create(ctx context.Context, user domain.Users) error
	GetByCredentials(ctx context.Context, name string) (domain.Users, error)
	GetPasswordHashByUsername(ctx context.Context, name string) (string, error)
	SetSession(ctx context.Context, session domain.Session, userID uuid.UUID) error
	GetUserIdByRefreshToken(ctx context.Context, refreshToken string) (uuid.UUID, error)
}

type Repositories struct {
	UsersRepo Users
}

func NewRepositories(conn *pgx.Conn) *Repositories {
	return &Repositories{
		UsersRepo: NewUsersRepo(conn),
	}
}
