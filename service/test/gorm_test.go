package test

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
	"time"
)

type Url struct {
	Id      int       `gorm:"column:id" json:"id"`
	Url     string    `gorm:"column:url" json:"url"`
	UrlName string    `gorm:"column:url_name" json:"urlName"`
	Remark  string    `gorm:"column:remark" json:"remark"`
	Date    time.Time `gorm:"column:date" json:"date"`
}

func (table *Url) TableName() string {
	return "url"
}

func TestGorm(t *testing.T) {
	userName := "root"
	password := "970617Baiyu!"
	dsn := userName + ":" + password + "@tcp(localhost:3306)/super_blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
		return
	}

	data := make([]*Url, 0)
	err = db.Debug().Find(&data).Error

	if err != nil {
		t.Fatal(err)
	}

	for _, datum := range data {
		fmt.Printf("Problem = %v \n", datum)
	}
}
