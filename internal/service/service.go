package service

import (
	"context"
	"log/slog"
	"tasktracker/internal/domain"
	repository "tasktracker/internal/repository/postgres"
)

type Deps struct {
	Repos *repository.Repositories
	Log   slog.Logger
}

type Users interface {
	SignUp(ctx context.Context, username string, password string) error
	SignIn(ctx context.Context, username string, password string) (domain.Tokens, error)
}

type Services struct {
	Users Users
}

func NewServices(deps Deps) *Services {
	return &Services{
		Users: NewUsersService(deps.Repos.UsersRepo, deps.Log),
	}
}
