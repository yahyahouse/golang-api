package provider

import (
	configImport "exam/config"
	"exam/internal/delivery/handler"
	"exam/internal/usecase"
)

var (
	calcGradeHandler  handler.HandlerFuncProvider
	getStudentHandler handler.HandlerFuncProvider
	gradingUseCase    usecase.GradingUseCase
)

func calcGradeHandlerFunc() (handler.HandlerFuncProvider, error) {
	if calcGradeHandler == nil {
		cfg := configImport.GetConfig()
		calcGradeHandler = handler.NewCalcGradeHandler(&cfg, getGradingUseCase())
	}
	return calcGradeHandler, nil
}

func getStudentHandlerFunc() (handler.HandlerFuncProvider, error) {
	if getStudentHandler == nil {
		cfg := configImport.GetConfig()
		getStudentHandler = handler.NewGetStudentHandler(&cfg, getGradingUseCase())

	}
	return getStudentHandler, nil
}

func getGradingUseCase() usecase.GradingUseCase {
	if gradingUseCase == nil {
		cfg := getConfig()

		gradingUseCase = usecase.NewGradingUseCase(&cfg, getRepositoryPostgres(), getRepositoryRedis())

	}
	return gradingUseCase
}
