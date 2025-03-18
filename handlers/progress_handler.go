package handlers

import (
	"net/http"

	"github.com/Alwin18/king-code/models"
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

func (h *ProgressHandler) UpdateUserProgress(c *gin.Context) {
	userID := c.Param("userID")
	var progress models.UserProgress
	if err := c.ShouldBindJSON(&progress); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	progress.UserID = userID

	if err := h.Service.UpdateProgress(&progress); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user progress"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User progress updated successfully"})
}
