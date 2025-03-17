package handlers

import (
	"net/http"

	"github.com/Alwin18/king-code/services"
	"github.com/gin-gonic/gin"
)

type LevelHandler struct {
	Service *services.LevelService
}

func NewLevelHandler(service *services.LevelService) *LevelHandler {
	return &LevelHandler{Service: service}
}

// Get All Levels
func (h *LevelHandler) GetAllLevels(c *gin.Context) {
	levels, err := h.Service.GetAllLevels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get levels"})
		return
	}
	c.JSON(http.StatusOK, levels)
}
