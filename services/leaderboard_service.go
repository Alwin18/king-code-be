package services

import (
	"github.com/Alwin18/king-code/models"
	"github.com/Alwin18/king-code/repositories"
)

type LeaderboardService struct {
	Repo *repositories.LeaderboardRepository
}

func NewLeaderboardService(repo *repositories.LeaderboardRepository) *LeaderboardService {
	return &LeaderboardService{Repo: repo}
}

func (s *LeaderboardService) GetTopPlayers(limit int) ([]models.Leaderboard, error) {
	return s.Repo.GetTopPlayers(limit)
}

func (s *LeaderboardService) GetLeaderboardByUserID(userID string) (*models.Leaderboard, error) {
	return s.Repo.GetLeaderboardByUserID(userID)
}

func (s *LeaderboardService) GetLeaderboard(limit int) ([]models.LeaderboardEntry, error) {
	return s.Repo.GetLeaderboard(limit)
}
