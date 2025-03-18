package handlers

import (
	"net/http"

	"github.com/Alwin18/king-code/models"
	"github.com/Alwin18/king-code/routes/response"
	"github.com/Alwin18/king-code/services"
	"github.com/gin-gonic/gin"
)

type LeaderboardHandler struct {
	Service *services.LeaderboardService
}

func NewLeaderboardHandler(service *services.LeaderboardService) *LeaderboardHandler {
	return &LeaderboardHandler{Service: service}
}

func (h *LeaderboardHandler) GetTopPlayers(c *gin.Context) {
	topPlayers, err := h.Service.GetTopPlayers(10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get top players"})
		return
	}
	c.JSON(http.StatusOK, topPlayers)
}

func (h *LeaderboardHandler) GetLeaderboardByUserID(c *gin.Context) {
	userID := c.Param("userID")
	leaderboard, err := h.Service.GetLeaderboardByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get leaderboard"})
		return
	}
	c.JSON(http.StatusOK, leaderboard)
}

// GetLeaderboard - Handler untuk mendapatkan leaderboard
func (h *LeaderboardHandler) GetLeaderboard(c *gin.Context) {
	leaderboard, err := h.Service.GetLeaderboard(10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch leaderboard"})
		return
	}

	c.JSON(http.StatusOK, response.Response[[]models.LeaderboardEntry]{Status: http.StatusOK, Message: "Success", Data: leaderboard})
}
