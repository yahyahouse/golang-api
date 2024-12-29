package handler

import (
	"exam/config"
	"exam/domain"
	"exam/internal/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CalcGradeHandler struct {
	cfg *config.Config
	usecase.GradingUseCase
}

func (h CalcGradeHandler) GetHandlerFunc() (method, path string, handler gin.HandlerFunc) {
	return post, "/student/calcgrade", h.handle
}

func (h CalcGradeHandler) handle(context *gin.Context) {
	var (
		req domain.Student
		res domain.Student
	)
	if err := context.BindJSON(&req); err != nil {
		// Handle error if binding fails
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx := context.Request.Context()
	res, _ = h.CalcGrade(ctx, req)
	context.JSON(http.StatusOK, res)

}

func NewCalcGradeHandler(cfg *config.Config, usecase usecase.GradingUseCase) CalcGradeHandler {
	return CalcGradeHandler{cfg, usecase}
}
