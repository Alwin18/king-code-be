package services

import (
	"github.com/Alwin18/king-code/models"
	"github.com/Alwin18/king-code/repositories"
)

type ChallengeService struct {
	Repo *repositories.ChallengeRepository
}

func NewChallengeService(repo *repositories.ChallengeRepository) *ChallengeService {
	return &ChallengeService{Repo: repo}
}

func (s *ChallengeService) CreateChallenge(challenge *models.CodingChallenge) error {
	return s.Repo.CreateChallenge(challenge)
}

func (s *ChallengeService) GetChallengesByLevel(levelID string) ([]models.CodingChallenge, error) {
	return s.Repo.GetChallengesByLevel(levelID)
}
