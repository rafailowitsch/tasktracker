package service

import (
	"context"
	"fmt"
	"log/slog"
	"tasktracker/internal/domain"
	repository "tasktracker/internal/repository/postgres"
	"tasktracker/pkg/auth"
	"tasktracker/pkg/log/sl"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UsersService struct {
	usersRepo repository.Users

	tokenManager auth.Manager
	log          slog.Logger
}

func NewUsersService(usersRepo repository.Users, tokenManager auth.Manager, log slog.Logger) *UsersService {
	return &UsersService{
		usersRepo:    usersRepo,
		tokenManager: tokenManager,
		log:          log,
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
	var (
		user   domain.Users
		err    error
		tokens Tokens
	)

	err = u.verifyPassword(ctx, name, password)
	if err != nil {
		u.log.Error("Failed verifying password")
		return Tokens{}, err
	}

	user, err = u.usersRepo.GetByCredentials(ctx, name)
	if err != nil {
		u.log.Error("credentials error", sl.Err(err))
		return Tokens{}, err
	} else {
		u.log.Info(fmt.Sprintf("successful sign-in %s", user.Name))
	}

	tokens, err = u.createSession(ctx, user.ID)
	if err != nil {
		u.log.Error("Failed creating session", sl.Err(err))
		return Tokens{}, err
	}

	return tokens, nil
}

func (u *UsersService) verifyPassword(ctx context.Context, name, password string) error {
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

func (u *UsersService) createSession(ctx context.Context, userID uuid.UUID) (Tokens, error) {
	u.log.Info("createSession")
	var (
		res Tokens
		err error
	)

	res.AccessToken, err = u.tokenManager.NewJWT(userID.String(), 2*time.Hour)
	if err != nil {
		return Tokens{}, err
	}
	// u.log.Info("access token initialized")

	res.RefreshToken, err = u.tokenManager.NewRefreshToken()
	if err != nil {
		return Tokens{}, err
	}
	// u.log.Info("refresh token initialized")

	session := domain.Session{
		RefreshToken: res.RefreshToken,
		ExpiresAt:    time.Now().Add(2 * time.Hour),
	}
	// u.log.Info("session initialized")

	err = u.usersRepo.SetSession(ctx, session, userID)
	if err != nil {
		return Tokens{}, err
	}
	// u.log.Info("successful set session")

	return res, nil
}
