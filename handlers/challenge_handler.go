package handlers

import (
	"net/http"

	"github.com/Alwin18/king-code/models"
	"github.com/Alwin18/king-code/routes/response"
	"github.com/Alwin18/king-code/services"
	"github.com/gin-gonic/gin"
)

type ChallengeHandler struct {
	Service *services.ChallengeService
}

func NewChallengeHandler(service *services.ChallengeService) *ChallengeHandler {
	return &ChallengeHandler{Service: service}
}

// GetChallengesByLevel - Handler untuk mendapatkan tantangan berdasarkan level
func (h *ChallengeHandler) GetChallengesByLevel(c *gin.Context) {
	levelID := c.Param("level_id")

	challenges, err := h.Service.GetChallengesByLevel(c, levelID)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, response.Response[[]models.Challenge]{Status: http.StatusOK, Message: "Success", Data: challenges})
}

// GetChallengeByID - Handler untuk mendapatkan detail tantangan
func (h *ChallengeHandler) GetChallengeByID(c *gin.Context) {
	challengeID := c.Param("id")

	challenge, err := h.Service.GetChallengeByID(c, challengeID)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{"challenge": challenge})
}
