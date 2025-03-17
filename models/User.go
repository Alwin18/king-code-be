package models

import "time"

type User struct {
	ID        string `gorm:"type:uuid;primaryKey"`
	Username  string `gorm:"unique"`
	Email     string `gorm:"unique"`
	Password  string
	XP        int
	Level     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (User) TableName() string {
	return "users"
}
