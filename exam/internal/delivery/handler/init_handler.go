/*
Package handler consist of http handler implementation.
Http handler logic only :
  - encode decode the request response
  - handle http response code and body based on usecase returned value/error

You should have no business logic handling in the implementation.
*/
package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

const (
	delete           = "DELETE"
	put              = "PUT"
	post             = "POST"
	get              = "GET"
	handlerStackName = "templateservice-bs.internal."
)

func init() {
	//setup gin log mode
	ginMode := viper.GetString("GIN_MODE")
	if ginMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

}

// HandlerFuncProvider serve as the interface that all concrete handler must implement
type HandlerFuncProvider interface {
	GetHandlerFunc() (method, path string, handler gin.HandlerFunc)
}

func GetPingHandler() (method, path string, handler gin.HandlerFunc) {
	return get, "/_internal/_ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"status": "OK"})
	}
}
