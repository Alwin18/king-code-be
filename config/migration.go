package config

import "github.com/Alwin18/king-code/models"

func MigrateDatabase() {
	DB.AutoMigrate(
		&models.User{},
		&models.Level{},
		&models.UserProgress{},
		&models.CodingChallenge{},
		&models.UserSubmission{},
		&models.Leaderboard{},
		&models.WebSocketLog{},
	)
}
