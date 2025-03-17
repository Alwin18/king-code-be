package models

import "time"

type WebSocketLog struct {
	ID        string `gorm:"type:uuid;primaryKey"`
	UserID    string `gorm:"type:uuid;not null"`
	EventType string `gorm:"not null"`
	Data      string `gorm:"type:jsonb"`
	Timestamp time.Time
}
