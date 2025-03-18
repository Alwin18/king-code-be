package handlers

import (
	"net/http"

	"github.com/Alwin18/king-code/models"
	"github.com/Alwin18/king-code/routes/response"
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
	c.JSON(http.StatusOK, response.Response[[]models.Level]{Status: http.StatusOK, Message: "Success", Data: levels})
}

func (h *LevelHandler) CreateLevel(c *gin.Context) {
	var level models.Level
	if err := c.ShouldBindJSON(&level); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if err := h.Service.CreateLevel(&level); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create level"})
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponse{Status: http.StatusCreated, Message: "Success Create Level"})
}
