package config

import (
	"github.com/Alwin18/king-code/models"
	"gorm.io/gorm"
)

func MigrateDatabase(DB *gorm.DB) {
	DB.AutoMigrate(
		&models.User{},
		&models.Level{},
		&models.UserProgress{},
		&models.Challenge{},
		&models.TestCase{},
		&models.UserSubmission{},
		&models.Leaderboard{},
		&models.WebSocketLog{},
		&models.RefreshToken{},
	)
}
