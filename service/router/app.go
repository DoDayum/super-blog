package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"super-blog/service"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/hello", func(context *gin.Context) {
		context.String(http.StatusOK, "hello")
	})

	r.GET("/queryUrl", service.QueryUrl)

	return r
}
