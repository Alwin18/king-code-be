package models

import "time"

type UserSubmission struct {
	ID          string `gorm:"type:uuid;primaryKey"`
	UserID      string `gorm:"type:uuid;not null"`
	ChallengeID string `gorm:"type:uuid;not null"`
	Code        string `gorm:"not null"`
	Status      string `gorm:"type:varchar(10);not null;check:status IN ('correct', 'incorrect, pending');default:'pending'"`
	SubmittedAt time.Time
}
