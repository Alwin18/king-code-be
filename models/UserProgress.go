package models

import (
	"time"

	"github.com/google/uuid"
)

type UserProgress struct {
	ID          string     `gorm:"type:uuid;primaryKey" json:"id"`
	UserID      string     `gorm:"type:uuid" json:"user_id"`
	LevelID     string     `gorm:"type:uuid" json:"level_id"`
	Status      string     `gorm:"type:varchar(10);not null;check:status IN ('pending', 'completed')" json:"status"`
	Score       int        `json:"score"`
	CompletedAt *time.Time `json:"completed_at"`
}

func (u *UserProgress) Default() {
	u.ID = uuid.New().String()
	u.CompletedAt = nil
	u.Status = "pending"
	u.Score = 0
}
