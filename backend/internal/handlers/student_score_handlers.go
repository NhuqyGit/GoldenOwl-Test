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

func (h *StudentScoreHandler) GetScoreBySBD(c *gin.Context) {
	sbd := c.Query("sbd")
	if sbd == "" {
		c.JSON(400, gin.H{"error": "Missing sbd query parameter"})
		return
	}

	score, err := h.service.GetScoreBySBD(sbd)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if score == nil {
		c.JSON(404, gin.H{"message": "Score not found"})
		return
	}
	c.JSON(200, score)
}

func (h *StudentScoreHandler) GetScoreReportBySubject(c *gin.Context){
	data, err := h.service.GetScoreReportBySubject()
	if err != nil{
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if data == nil{
		c.JSON(404, gin.H{"message": "Data not found"})
		return
	}

	c.JSON(200, data)
}

func (h *StudentScoreHandler) GetTop10GroupA(c *gin.Context){
	scores, err := h.service.GetTop10GroupA()
	if err != nil{
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	if scores == nil{
		c.JSON(404, gin.H{"message": "Data not found"})
		return
	}

	c.JSON(200, scores)
}
