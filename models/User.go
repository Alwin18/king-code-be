package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

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

func (u *User) Validate() error {
	if u.Username == "" || u.Email == "" || u.Password == "" {
		return errors.New("username, email, and password are required")
	}
	return nil
}

func (u *User) Default() {
	u.ID = uuid.New().String()
	u.CreatedAt = time.Now()
	u.XP = 0
	u.Level = 1
}
