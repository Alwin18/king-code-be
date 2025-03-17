package models

import "time"

type Level struct {
	ID          string `gorm:"type:uuid;primaryKey"`
	Title       string
	Description string
	Difficulty  string `gorm:"type:enum('easy', 'medium', 'hard')"`
	XPReward    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
