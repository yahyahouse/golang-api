package handler

import (
	"exam/config"
	"exam/domain"
	"exam/internal/usecase"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type GetStudentHandler struct {
	cfg *config.Config
	usecase.GradingUseCase
}

func (h GetStudentHandler) GetHandlerFunc() (method, path string, handler gin.HandlerFunc) {
	return get, "/student", h.handle
}

func (h GetStudentHandler) handle(context *gin.Context) {
	var (
		req string
		res domain.Student
	)
	//if err := context.Query(req);  {
	//	// Handle error if binding fails
	//	context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//	return
	//}
	req = context.Query("name")
	log.Println(req)
	ctx := context.Request.Context()
	res, _ = h.GetStudent(ctx, req)
	log.Println(res)
	context.JSON(http.StatusOK, res)

}

func NewGetStudentHandler(cfg *config.Config, usecase usecase.GradingUseCase) GetStudentHandler {
	return GetStudentHandler{cfg, usecase}
}
