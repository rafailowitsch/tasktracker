package repository

import (
	"context"
	"tasktracker/internal/domain"

	"github.com/jackc/pgx/v5"
)

type Users interface {
	Create(ctx context.Context, user domain.Users) error
	GetByCredentials(ctx context.Context, username, password string) (domain.Users, error)
}

// type Deps struct {
// 	DB  *pgx.Conn
// 	log slog.Logger
// }

type Repositories struct {
	UsersRepo Users
}

func NewRepositories(conn *pgx.Conn) *Repositories {
	return &Repositories{
		UsersRepo: NewUsersRepo(conn),
	}
}
