package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TestCase struct {
	ID             string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	ChallengeID    string `gorm:"type:uuid;not null"`
	Input          string `gorm:"type:text;not null"`
	ExpectedOutput string `gorm:"type:text;not null"`
}

func (tc *TestCase) BeforeCreate(tx *gorm.DB) error {
	tc.ID = uuid.New().String()
	return nil
}
