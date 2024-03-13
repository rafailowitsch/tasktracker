package domain

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	PasswordHash string    `json:"password_hash"`
}

type Session struct {
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}
