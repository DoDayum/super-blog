package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB = Init()

func Init() *gorm.DB {

	fmt.Printf("数据库初始化了")

	userName := "root"
	password := "970617Baiyu!"
	database := "super_blog"
	ip := "localhost:3306"

	dsn := userName + ":" + password + "@tcp(" + ip + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
		//log.Println("gorm init error:", err)
	}

	return db

}
