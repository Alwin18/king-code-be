package models

import (
	"time"

	"github.com/google/uuid"
)

type Level struct {
	ID          string `gorm:"type:uuid;primaryKey" json:"id"`
	Title       string `gorm:"not null" json:"title"`
	Description string `json:"description"`
	Difficulty  string `gorm:"type:varchar(10);not null;check:difficulty IN ('easy', 'medium', 'hard')" json:"difficulty"`
	XPReward    int    `gorm:"not null" json:"xp_reward"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (u *Level) Default() {
	u.ID = uuid.New().String()
	u.CreatedAt = time.Now()
}
