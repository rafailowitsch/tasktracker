package service

import (
	"context"
	"fmt"
	"log/slog"
	"tasktracker/internal/domain"
	repository "tasktracker/internal/repository/postgres"
	"tasktracker/pkg/hasher"
	"tasktracker/pkg/log/sl"
)

type UsersService struct {
	usersRepo repository.Users

	log slog.Logger
}

func NewUsersService(usersRepo repository.Users, log slog.Logger) *UsersService {
	return &UsersService{
		usersRepo: usersRepo,
		log:       log,
	}
}

func (u *UsersService) SignUp(ctx context.Context, name string, password string) error {
	password_hash := hasher.Hash(password)

	user := domain.Users{
		Name:     name,
		Password: password_hash,
	}

	if err := u.usersRepo.Create(ctx, user); err != nil {
		u.log.Error("failed creating user", sl.Err(err))
	} else {
		u.log.Info(fmt.Sprintf("user %s create", user.Name))
	}

	return nil
}

func (u *UsersService) SignIn(ctx context.Context, username string, password string) (Tokens, error) {
	token := Tokens{}
	return token, nil
}
