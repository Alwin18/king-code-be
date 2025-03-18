package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWT_SECRET = []byte("secret_key")

type JWTClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

// Generate Access Token (15 menit)
func GenerateAccessToken(userID string) (string, error) {
	claims := JWTClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWT_SECRET)
}

// Generate Refresh Token (7 hari)
func GenerateRefreshToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	return token.SignedString(JWT_SECRET)
}

// Parse Token
func ParseToken(tokenStr string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JWT_SECRET, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
