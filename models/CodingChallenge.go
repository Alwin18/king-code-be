package models

import "time"

type CodingChallenge struct {
	ID             string `gorm:"type:uuid;primaryKey"`
	LevelID        string `gorm:"type:uuid;not null"`
	Question       string `gorm:"not null"`
	ExpectedOutput string `gorm:"not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
