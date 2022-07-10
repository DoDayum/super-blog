package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/hello", func(context *gin.Context) {
		context.String(http.StatusOK, "hello")
	})

	return r
}
