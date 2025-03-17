package repositories

import (
	"github.com/Alwin18/king-code/models"
	"gorm.io/gorm"
)

type SubmissionRepository struct {
	DB *gorm.DB
}

func NewSubmissionRepository(db *gorm.DB) *SubmissionRepository {
	return &SubmissionRepository{DB: db}
}

func (r *SubmissionRepository) CreateSubmission(submission *models.UserSubmission) error {
	return r.DB.Create(submission).Error
}

func (r *SubmissionRepository) GetSubmissionsByUser(userID string) ([]models.UserSubmission, error) {
	var submissions []models.UserSubmission
	result := r.DB.Where("user_id = ?", userID).Find(&submissions)
	return submissions, result.Error
}
