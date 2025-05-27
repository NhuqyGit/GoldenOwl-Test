package routers

import (
	"goldenowl-test/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterSubjectRoutes(r *gin.Engine, handler *handlers.StudentScoreHandler) {
	subjects := r.Group("/student-scores")
	{
		subjects.GET("/", handler.GetFirst10Rows)
	}
}
