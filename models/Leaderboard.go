package models

import "time"

type Leaderboard struct {
	ID        string `gorm:"type:uuid;primaryKey"`
	UserID    string `gorm:"type:uuid;not null"`
	TotalXP   int    `gorm:"not null"`
	Rank      int
	UpdatedAt time.Time
}

type LeaderboardEntry struct {
	ID         string `json:"id"`
	Username   string `json:"username"`
	TotalScore int    `json:"total_score"`
}
