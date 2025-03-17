package repositories

import (
	"github.com/Alwin18/king-code/models"
	"gorm.io/gorm"
)

type ChallengeRepository struct {
	DB *gorm.DB
}

func (r *ChallengeRepository) CreateChallenge(challenge *models.CodingChallenge) error {
	return r.DB.Create(challenge).Error
}

func (r *ChallengeRepository) GetChallengesByLevel(levelID string) ([]models.CodingChallenge, error) {
	var challenges []models.CodingChallenge
	result := r.DB.Where("level_id = ?", levelID).Find(&challenges)
	return challenges, result.Error
}
