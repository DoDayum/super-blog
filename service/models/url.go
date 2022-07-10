package models

import "time"

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
