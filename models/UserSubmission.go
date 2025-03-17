package models

import "time"

type UserSubmission struct {
	ID          string `gorm:"type:uuid;primaryKey"`
	UserID      string `gorm:"type:uuid;not null"`
	ChallengeID string `gorm:"type:uuid;not null"`
	Code        string `gorm:"not null"`
	Status      string `gorm:"type:enum('pending', 'correct', 'incorrect');default:'pending'"`
	SubmittedAt time.Time
}
