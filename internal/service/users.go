package service

import (
	"context"
	"fmt"
	"log/slog"
	"tasktracker/internal/domain"
	repository "tasktracker/internal/repository/postgres"
	"tasktracker/pkg/log/sl"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

func (u *UsersService) SignUp(ctx context.Context, name, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		u.log.Error("error hasing: ", sl.Err(err))
	}
	password_hash := string(hash)

	user := domain.Users{
		ID:           uuid.New(),
		Name:         name,
		PasswordHash: password_hash,
	}

	if err := u.usersRepo.Create(ctx, user); err != nil {
		u.log.Error("failed creating user", sl.Err(err))
	} else {
		u.log.Info(fmt.Sprintf("user %s create with id %s", user.Name, user.ID.String()))
	}

	return nil
}

func (u *UsersService) SignIn(ctx context.Context, name, password string) (Tokens, error) {
	if err := u.VerifyPassword(ctx, name, password); err != nil {
		u.log.Error("Failed verifying password")
		return Tokens{}, err
	}

	if user, err := u.usersRepo.GetByCredentials(ctx, name); err != nil {
		u.log.Error("credentials error", sl.Err(err))
	} else {
		u.log.Info(fmt.Sprintf("successful sign-in %s", user.Name))
	}

	token := Tokens{}
	return token, nil
}

func (u *UsersService) VerifyPassword(ctx context.Context, name, password string) error {
	password_hash, err := u.usersRepo.GetPasswordHashByUsername(ctx, name)
	if err != nil {
		u.log.Error("Failed password hash retrieval", sl.Err(err))
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(password_hash), []byte(password))
	if err != nil {
		u.log.Error("Wrong password", sl.Err(err))
		return err
	}

	return nil
}
