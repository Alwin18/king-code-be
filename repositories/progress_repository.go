package repositories

import (
	"github.com/Alwin18/king-code/models"
	"gorm.io/gorm"
)

type ProgressRepository struct {
	DB *gorm.DB
}

func (r *ProgressRepository) CreateProgress(progress *models.UserProgress) error {
	return r.DB.Create(progress).Error
}

func (r *ProgressRepository) GetUserProgress(userID string) ([]models.UserProgress, error) {
	var progress []models.UserProgress
	result := r.DB.Where("user_id = ?", userID).Find(&progress)
	return progress, result.Error
}

func (r *ProgressRepository) UpdateProgress(progress *models.UserProgress) error {
	return r.DB.Save(progress).Error
}
