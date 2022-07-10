package service

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"super-blog/models"
)

func QueryUrl(c *gin.Context) {
	data := make([]*models.Url, 0)
	err := models.DB.Debug().Find(&data).Error
	if err != nil {
		log.Println("query has error:")
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
	})
}
