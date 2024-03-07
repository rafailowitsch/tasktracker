package domain

import "github.com/google/uuid"

type Users struct {
	ID           uuid.UUID `json:"id" db: "id"`
	Name         string    `json:"name" db: "name"`
	PasswordHash string    `json:"password_hash" db: "password_hash"`
}
