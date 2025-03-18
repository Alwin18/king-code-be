package handlers

import (
	"net/http"
	"time"

	"github.com/Alwin18/king-code/entity"
	"github.com/Alwin18/king-code/models"
	"github.com/Alwin18/king-code/repositories"
	"github.com/Alwin18/king-code/routes/response"
	"github.com/Alwin18/king-code/services"
	"github.com/Alwin18/king-code/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userService *services.UserService
	tokenRepo   *repositories.TokenRepository
}

func NewAuthHandler(userService *services.UserService, tokenRepo *repositories.TokenRepository) *AuthHandler {
	return &AuthHandler{userService, tokenRepo}
}

// Login Handler
func (h *AuthHandler) Login(c *gin.Context) {
	var request entity.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	user, err := h.userService.AuthenticateUser(request.Email, request.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate Tokens
	accessToken, _ := utils.GenerateAccessToken(user.ID)
	refreshToken, _ := utils.GenerateRefreshToken()
	h.tokenRepo.DeleteRefreshToken(user.ID)

	// Simpan refresh token ke database
	refreshTokenModel := models.NewRefreshToken(user.ID, refreshToken, 7*24*time.Hour)

	h.tokenRepo.SaveRefreshToken(refreshTokenModel)

	c.JSON(http.StatusOK, response.Response[entity.LoginResponse]{Status: http.StatusOK, Message: "Success", Data: entity.LoginResponse{
		GetUserByIdResponse: *user,
		Token:               accessToken,
		RefreshToken:        refreshToken,
	}})
}
