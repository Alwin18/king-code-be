package repositories

import (
	"github.com/Alwin18/king-code/models"
	"gorm.io/gorm"
)

type LeaderboardRepository struct {
	DB *gorm.DB
}

func (r *LeaderboardRepository) GetTopPlayers(limit int) ([]models.Leaderboard, error) {
	var leaderboard []models.Leaderboard
	result := r.DB.Order("xp DESC").Limit(limit).Find(&leaderboard)
	return leaderboard, result.Error
}

func (r *LeaderboardRepository) GetLeaderboardByUserID(userID string) (*models.Leaderboard, error) {
	var leaderboard models.Leaderboard
	result := r.DB.First(&leaderboard, "user_id = ?", userID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &leaderboard, nil
}
