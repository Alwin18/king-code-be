package services

import (
	"github.com/Alwin18/king-code/models"
	"github.com/Alwin18/king-code/repositories"
)

type ProgressService struct {
	Repo *repositories.ProgressRepository
}

func NewProgressService(repo *repositories.ProgressRepository) *ProgressService {
	return &ProgressService{Repo: repo}
}

func (s *ProgressService) CreateProgress(progress *models.UserProgress) error {
	return s.Repo.CreateProgress(progress)
}

func (s *ProgressService) GetUserProgress(userID string) ([]models.UserProgress, error) {
	return s.Repo.GetUserProgress(userID)
}

func (s *ProgressService) UpdateProgress(progress *models.UserProgress) error {
	return s.Repo.UpdateProgress(progress)
}
