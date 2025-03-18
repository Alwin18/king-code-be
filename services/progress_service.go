package services

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Alwin18/king-code/entity"
	"github.com/Alwin18/king-code/models"
	"github.com/Alwin18/king-code/repositories"
	"github.com/Alwin18/king-code/routes/response"
	"github.com/gin-gonic/gin"
)

type ProgressService struct {
	Repo      *repositories.ProgressRepository
	LevelRepo *repositories.LevelRepository
}

func NewProgressService(repo *repositories.ProgressRepository, levelRepo *repositories.LevelRepository) *ProgressService {
	return &ProgressService{Repo: repo, LevelRepo: levelRepo}
}

func (s *ProgressService) CreateProgress(progress *models.UserProgress) error {
	progress.Default()
	fmt.Println(progress)
	return s.Repo.CreateProgress(progress)
}

func (s *ProgressService) GetUserProgress(userID string) ([]models.UserProgress, error) {
	return s.Repo.GetUserProgress(userID)
}

func (s *ProgressService) UpdateProgress(c *gin.Context, body *entity.UpdateProgressRequest) error {
	level, err := s.LevelRepo.GetLevelByID(body.LevelID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
		return err
	}
	if level == nil {
		c.JSON(http.StatusInternalServerError, response.ErrorResponse{Status: http.StatusInternalServerError, Message: "Level not found"})
		return errors.New("Level not found")
	}

	now := time.Now()
	progress := &models.UserProgress{
		ID:          body.ProgressID,
		UserID:      body.UserID,
		LevelID:     body.LevelID,
		Status:      "completed",
		Score:       level.XPReward,
		CompletedAt: &now,
	}

	return s.Repo.UpdateProgress(progress)
}
