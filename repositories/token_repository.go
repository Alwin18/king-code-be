package repositories

import (
	"github.com/Alwin18/king-code/models"
	"gorm.io/gorm"
)

type TokenRepository struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) *TokenRepository {
	return &TokenRepository{db}
}

// Simpan Refresh Token
func (r *TokenRepository) SaveRefreshToken(token *models.RefreshToken) error {
	return r.db.Create(token).Error
}

// Ambil Refresh Token
func (r *TokenRepository) GetRefreshToken(token string) (*models.RefreshToken, error) {
	var refreshToken models.RefreshToken
	err := r.db.Where("token = ?", token).First(&refreshToken).Error
	if err != nil {
		return nil, err
	}
	return &refreshToken, nil
}

// Hapus Refresh Token (Saat Logout)
func (r *TokenRepository) DeleteRefreshToken(userId string) error {
	return r.db.Where("user_id = ?", userId).Delete(&models.RefreshToken{}).Error
}
