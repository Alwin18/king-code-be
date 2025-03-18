package models

import (
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID    string    `gorm:"type:uuid;not null"`
	Token     string    `gorm:"unique;not null"`
	ExpiresAt time.Time `gorm:"not null"`
	CreatedAt time.Time
}

// Generate new refresh token
func NewRefreshToken(userID string, token string, duration time.Duration) *RefreshToken {
	return &RefreshToken{
		ID:        uuid.New(),
		UserID:    userID,
		Token:     token,
		ExpiresAt: time.Now().Add(duration),
	}
}
