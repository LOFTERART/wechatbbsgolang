package models

import "github.com/jinzhu/gorm"

type SubTopic struct {
	gorm.Model
	Name string
	SubName string
	SendNum int   //发送帖子数
	Follow int    //关注人数
	ClassifyId int
}