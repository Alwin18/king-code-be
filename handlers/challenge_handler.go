package handlers

import (
	"net/http"

	"github.com/Alwin18/king-code/services"
	"github.com/gin-gonic/gin"
)

type ChallengeHandler struct {
	Service *services.ChallengeService
}

func NewChallengeHandler(service *services.ChallengeService) *ChallengeHandler {
	return &ChallengeHandler{Service: service}
}

// Get Challenges by Level
func (h *ChallengeHandler) GetChallengesByLevel(c *gin.Context) {
	levelID := c.Param("levelID")
	challenges, err := h.Service.GetChallengesByLevel(levelID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get challenges"})
		return
	}
	c.JSON(http.StatusOK, challenges)
}
