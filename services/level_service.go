package services

import (
	"github.com/Alwin18/king-code/models"
	"github.com/Alwin18/king-code/repositories"
)

type LevelService struct {
	Repo *repositories.LevelRepository
}

func NewLevelService(repo *repositories.LevelRepository) *LevelService {
	return &LevelService{Repo: repo}
}

func (s *LevelService) CreateLevel(level *models.Level) error {
	level.Default()
	return s.Repo.CreateLevel(level)
}

func (s *LevelService) GetAllLevels() ([]models.Level, error) {
	return s.Repo.GetAllLevels()
}

func (s *LevelService) GetLevelByID(id string) (*models.Level, error) {
	return s.Repo.GetLevelByID(id)
}
