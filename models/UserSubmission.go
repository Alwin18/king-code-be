package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserSubmission struct {
	ID          string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID      string    `gorm:"type:uuid;not null"`
	ChallengeID string    `gorm:"type:uuid;not null"`
	Code        string    `gorm:"type:text;not null"`
	Status      string    `gorm:"type:text;not null"` // "pending", "correct", "wrong"
	Score       int       `gorm:"type:int"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}

func (s *UserSubmission) BeforeCreate(tx *gorm.DB) error {
	s.ID = uuid.New().String()
	s.CreatedAt = time.Now()
	return nil
}
