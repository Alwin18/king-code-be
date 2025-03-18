package handlers

import (
	"net/http"

	"github.com/Alwin18/king-code/entity"
	"github.com/Alwin18/king-code/models"
	"github.com/Alwin18/king-code/routes/response"
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
	var progress entity.UpdateProgressRequest
	if err := c.ShouldBindJSON(&progress); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := h.Service.UpdateProgress(c, &progress); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user progress"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User progress updated successfully"})
}

func (h *ProgressHandler) CreateProgress(c *gin.Context) {
	var progress models.UserProgress
	if err := c.ShouldBindJSON(&progress); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if err := h.Service.CreateProgress(&progress); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create progress"})
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse{Status: http.StatusCreated, Message: "Success Create Progress"})
}
