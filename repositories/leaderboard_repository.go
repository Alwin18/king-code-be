package repositories

import (
	"github.com/Alwin18/king-code/models"
	"gorm.io/gorm"
)

type LeaderboardRepository struct {
	DB *gorm.DB
}

func NewLeaderboardRepository(db *gorm.DB) *LeaderboardRepository {
	return &LeaderboardRepository{DB: db}
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

// GetLeaderboard - Ambil daftar pemain berdasarkan skor tertinggi
func (r *LeaderboardRepository) GetLeaderboard(limit int) ([]models.LeaderboardEntry, error) {
	var leaderboard []models.LeaderboardEntry
	err := r.DB.Raw(`
		SELECT users.id, users.username, COALESCE(SUM(user_progresses.score), 0) AS total_score
		FROM users
		LEFT JOIN user_progresses ON users.id = user_progresses.user_id
		GROUP BY users.id
		ORDER BY total_score DESC
		LIMIT ?
	`, limit).Scan(&leaderboard).Error

	return leaderboard, err
}
