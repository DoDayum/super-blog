package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"log"
	"net/http"
	"super-blog/models"
)

func QueryUrl(c *gin.Context) {
	// x-www-form-urlencoded
	age := c.PostForm("age")
	fmt.Printf(age)

	// application/json
	body, _ := c.GetRawData()
	var m map[string]interface{}
	err := json.Unmarshal(body, &m)
	if err != nil {
		return
	}
	fmt.Printf("%v", m)

	data := make([]*models.Url, 0)
	err = models.DB.Debug().Find(&data).Error
	if err != nil {
		log.Println("query has error:")
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": data,
	})
}
