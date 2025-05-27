package handlers

import (
	"goldenowl-test/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StudentScoreHandler struct {
	service *services.StudentScoreService
}

func NewStudentScoreHandler(service *services.StudentScoreService) *StudentScoreHandler {
	return &StudentScoreHandler{service: service}
}

func (h *StudentScoreHandler) GetFirst10Rows(c *gin.Context) {
	StudentScores, err := h.service.GetFirst10Rows()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch StudentScores"})
		return
	}
	c.JSON(http.StatusOK, StudentScores)
}