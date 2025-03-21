package services

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Alwin18/king-code/models"
	"github.com/Alwin18/king-code/repositories"
	"github.com/Alwin18/king-code/routes/response"
	"github.com/gin-gonic/gin"
)

type ChallengeService struct {
	Repo *repositories.ChallengeRepository
}

func NewChallengeService(repo *repositories.ChallengeRepository) *ChallengeService {
	return &ChallengeService{Repo: repo}
}

// GetChallengesByLevel - Ambil daftar tantangan berdasarkan level
func (s *ChallengeService) GetChallengesByLevel(c *gin.Context, levelID string) ([]models.Challenge, error) {
	ressponse, err := s.Repo.GetChallengesByLevel(levelID)
	if err != nil {
		fmt.Println(err)
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, response.ErrorResponse{Status: http.StatusNotFound, Message: "Level not found"})
			return nil, errors.New("data not found")
		} else {
			c.JSON(http.StatusInternalServerError, response.ErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
			return nil, err
		}
	}
	return ressponse, err
}

// GetChallengeByID - Ambil detail tantangan berdasarkan ID
func (s *ChallengeService) GetChallengeByID(c *gin.Context, id string) (*models.Challenge, error) {
	return s.Repo.GetChallengeByID(id)
}
