package entity

import "time"

type GetUserByIdResponse struct {
	ID        string    `json:"ID"`
	Username  string    `json:"Username"`
	Email     string    `json:"Email"`
	XP        int       `json:"XP"`
	Level     int       `json:"Level"`
	CreatedAt time.Time `json:"CreatedAt"`
	UpdatedAt time.Time `json:"UpdatedAt"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	GetUserByIdResponse
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
