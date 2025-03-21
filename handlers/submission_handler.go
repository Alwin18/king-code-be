package handlers

import (
	"net/http"

	"github.com/Alwin18/king-code/services"
	"github.com/gin-gonic/gin"
)

type SubmissionHandler struct {
	service *services.SubmissionService
}

func NewSubmissionHandler(service *services.SubmissionService) *SubmissionHandler {
	return &SubmissionHandler{service}
}

// Submit jawaban user
func (h *SubmissionHandler) SubmitCode(c *gin.Context) {
	var req struct {
		UserID      string `json:"user_id"`
		ChallengeID string `json:"challenge_id"`
		Code        string `json:"code"`
		Language    string `json:"language"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	submission, err := h.service.EvaluateSubmission(req.UserID, req.ChallengeID, req.Code, req.Language)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": submission.Status,
		"score":  submission.Score,
	})
}
