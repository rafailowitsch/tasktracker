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
	RefreshToken string    `json:"refreshToken"`
	ExpiresAt    time.Time `json:"expiresAt"`
}
