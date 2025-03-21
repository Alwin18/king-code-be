package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Challenge struct {
	ID          string     `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	LevelID     string     `gorm:"type:uuid;not null"`
	Title       string     `gorm:"type:text;not null"`
	Description string     `gorm:"type:text;not null"`
	Language    string     `gorm:"type:text;not null"`
	TestCases   []TestCase `gorm:"foreignKey:ChallengeID"`
	CreatedAt   time.Time  `gorm:"autoCreateTime"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime"`
}

// Sebelum simpan, pastikan ID di-generate
func (c *Challenge) BeforeCreate(tx *gorm.DB) error {
	c.ID = uuid.New().String()
	c.CreatedAt = time.Now()
	return nil
}
