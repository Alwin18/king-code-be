package repositories

import (
	"github.com/Alwin18/king-code/models"
	"gorm.io/gorm"
)

type LevelRepository struct {
	DB *gorm.DB
}

func NewLevelRepository(db *gorm.DB) *LevelRepository {
	return &LevelRepository{DB: db}
}

func (r *LevelRepository) CreateLevel(level *models.Level) error {
	return r.DB.Create(level).Error
}

func (r *LevelRepository) GetAllLevels() ([]models.Level, error) {
	var levels []models.Level
	result := r.DB.Find(&levels)
	return levels, result.Error
}

func (r *LevelRepository) GetLevelByID(id string) (*models.Level, error) {
	var level models.Level
	result := r.DB.First(&level, "id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &level, nil
}
