package service

import (
	"context"
	"log/slog"
	repository "tasktracker/internal/repository/postgres"
)

type Deps struct {
	Repos *repository.Repositories
	Log   slog.Logger
}

type Users interface {
	SignUp(ctx context.Context, name string, password string) error
	SignIn(ctx context.Context, name string, password string) (Tokens, error)
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

type Services struct {
	Users Users
}

func NewServices(deps Deps) *Services {
	return &Services{
		Users: NewUsersService(deps.Repos.UsersRepo, deps.Log),
	}
}
