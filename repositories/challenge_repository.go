package repositories

import (
	"github.com/Alwin18/king-code/models"
	"gorm.io/gorm"
)

type ChallengeRepository struct {
	DB *gorm.DB
}

func NewChallengeRepository(db *gorm.DB) *ChallengeRepository {
	return &ChallengeRepository{DB: db}
}

// GetChallengesByLevel - Ambil semua tantangan berdasarkan level
func (r *ChallengeRepository) GetChallengesByLevel(levelID string) ([]models.Challenge, error) {
	var challenges []models.Challenge
	err := r.DB.Where("level_id = ?", levelID).Preload("TestCases").Find(&challenges).Error
	return challenges, err
}

// GetChallengeByID - Ambil tantangan berdasarkan ID
func (r *ChallengeRepository) GetChallengeByID(id string) (*models.Challenge, error) {
	var challenge models.Challenge
	err := r.DB.Where("id = ?", id).Preload("TestCases").First(&challenge).Error
	if err != nil {
		return nil, err
	}
	return &challenge, nil
}
