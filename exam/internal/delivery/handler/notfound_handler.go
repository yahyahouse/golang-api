package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, "Not Found")
}
