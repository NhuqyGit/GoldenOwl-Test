package handlers

import (
	"goldenowl-test/internal/services"

	"github.com/gin-gonic/gin"
)

type StudentScoreHandler struct {
	service *services.StudentScoreService
}

func NewStudentScoreHandler(service *services.StudentScoreService) *StudentScoreHandler {
	return &StudentScoreHandler{service: service}
}

// GetScoreBySBD godoc
// @Summary      Get student score by SBD
// @Description  Returns student score based on registration number (SBD)
// @Tags         student-scores
// @Produce      json
// @Param        sbd  query     string  true  "Student Registration Number"
// @Success      200  {object}  models.StudentScore
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /student-scores [get]
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

// GetScoreReportBySubject godoc
// @Summary      Get score report by subject
// @Description  Returns an aggregated report of scores by subject
// @Tags         student-scores
// @Produce      json
// @Success      200  {array}  models.StudentScore
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /student-scores/report [get]
func (h *StudentScoreHandler) GetScoreReportBySubject(c *gin.Context) {
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

// GetTop10GroupA godoc
// @Summary      Get top 10 scores for Group A
// @Description  Returns the top 10 students in Group A based on their scores
// @Tags         student-scores
// @Produce      json
// @Success      200  {array}  models.StudentScore
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /student-scores/top-10-groupA [get]
func (h *StudentScoreHandler) GetTop10GroupA(c *gin.Context) {
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
