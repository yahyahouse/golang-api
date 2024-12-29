package provider

import (
	"exam/internal/delivery/handler"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var (
	//mapRestPath will hold rest path that already served
	mapRestPath = map[string]bool{}
	//response code in this slice will not give error metrics
	//use slice instead of map for performance reason
	businessOkRC = []bool{
		http.StatusOK:                  true,
		http.StatusBadRequest:          true,
		http.StatusUnauthorized:        false,
		http.StatusNotFound:            false,
		http.StatusBadGateway:          false,
		http.StatusServiceUnavailable:  false,
		http.StatusInternalServerError: false,
		http.StatusGatewayTimeout:      false,
	}
	server *http.Server
)

type getHandlerFunc func() (handler.HandlerFuncProvider, error)

func getServerHandler() http.Handler {
	router := gin.New()

	//add internal handler that doesn't need to be wrapped with metrics logger
	router.Handle(handler.GetPingHandler())

	funcs := []getHandlerFunc{getStudentHandlerFunc, calcGradeHandlerFunc}
	putHandlersToRouter(router, funcs)
	router.NoRoute(handler.NotFound)
	return router
}
func putHandlersToRouter(router *gin.Engine, funcs []getHandlerFunc) {
	for _, theFunc := range funcs {
		handler, err := theFunc()
		if err != nil {
			panic(fmt.Sprintf("[Main] can't provide Handler %v", err))
		}
		method, path, handlerFunc := handler.GetHandlerFunc()
		router.Handle(method, path, handlerFunc)
	}
}
func (p providerImpl) GetServer() *http.Server {
	if server != nil {
		return server
	}

	serverConf := getConfig().ServerCfg
	readTimeout, err := time.ParseDuration(serverConf.ReadTimeout)
	if err != nil {
		panic(fmt.Sprintf("failed parse server read timeout %v", err))
	}

	writeTimeout, err := time.ParseDuration(serverConf.WriteTimeout)
	if err != nil {
		panic(fmt.Sprintf("failed parse server write timeout %v", err))
	}

	server = &http.Server{
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		Addr:         ":8083",
		Handler:      getServerHandler(),
	}
	return server
}
