package handlers

import (
	"net/http"

	"github.com/Alwin18/king-code/services"
	"github.com/gin-gonic/gin"
)

type ProgressHandler struct {
	Service *services.ProgressService
}

func NewProgressHandler(service *services.ProgressService) *ProgressHandler {
	return &ProgressHandler{Service: service}
}

// Get User Progress
func (h *ProgressHandler) GetUserProgress(c *gin.Context) {
	userID := c.Param("userID")
	progress, err := h.Service.GetUserProgress(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user progress"})
		return
	}
	c.JSON(http.StatusOK, progress)
}
