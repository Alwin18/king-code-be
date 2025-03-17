package handlers

import (
	"net/http"

	"github.com/Alwin18/king-code/models"
	"github.com/Alwin18/king-code/services"
	"github.com/gin-gonic/gin"
)

type SubmissionHandler struct {
	Service *services.SubmissionService
}

func NewSubmissionHandler(service *services.SubmissionService) *SubmissionHandler {
	return &SubmissionHandler{Service: service}
}

func (s *SubmissionHandler) CreateSubmission(c *gin.Context) {
	var submission models.UserSubmission
	if err := c.ShouldBindJSON(&submission); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if err := s.Service.CreateSubmission(&submission); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create submission"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Submission created successfully"})
}

func (s *SubmissionHandler) GetSubmissionsByUser(c *gin.Context) {
	userID := c.Param("userID")
	submissions, err := s.Service.GetSubmissionsByUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get submissions"})
		return
	}
	c.JSON(http.StatusOK, submissions)
}
