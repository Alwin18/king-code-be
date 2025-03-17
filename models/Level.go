package models

import "time"

type Level struct {
	ID          string `gorm:"type:uuid;primaryKey"`
	Title       string `gorm:"not null"`
	Description string
	Difficulty  string `gorm:"type:varchar(10);not null;check:difficulty IN ('easy', 'medium', 'hard')"`
	XPReward    int    `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
