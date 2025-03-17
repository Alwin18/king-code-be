package models

import "time"

type UserProgress struct {
	ID          string `gorm:"type:uuid;primaryKey"`
	UserID      string `gorm:"type:uuid"`
	LevelID     string `gorm:"type:uuid"`
	Status      string `gorm:"type:varchar(10);not null;check:status IN ('pending', 'completed')"`
	Score       int
	CompletedAt *time.Time
}
