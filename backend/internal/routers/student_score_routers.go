package routers

import (
	"goldenowl-test/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterSubjectRoutes(r *gin.Engine, handler *handlers.StudentScoreHandler) {
	subjects := r.Group("/student-scores")
	{
		subjects.GET("/", handler.GetScoreBySBD)
		subjects.GET("/report", handler.GetScoreReportBySubject)
		subjects.GET("/top-10-groupA", handler.GetTop10GroupA)
	}
}
