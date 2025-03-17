package services

import (
	"github.com/Alwin18/king-code/models"
	"github.com/Alwin18/king-code/repositories"
)

type SubmissionService struct {
	Repo *repositories.SubmissionRepository
}

func NewSubmissionService(repo *repositories.SubmissionRepository) *SubmissionService {
	return &SubmissionService{Repo: repo}
}

func (s *SubmissionService) CreateSubmission(submission *models.UserSubmission) error {
	return s.Repo.CreateSubmission(submission)
}

func (s *SubmissionService) GetSubmissionsByUser(userID string) ([]models.UserSubmission, error) {
	return s.Repo.GetSubmissionsByUser(userID)
}
