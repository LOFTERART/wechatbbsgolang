package models

import "github.com/jinzhu/gorm"

type SubTopic struct {
	gorm.Model
	Name string
	SubName string
	SendNum uint   //发送帖子数
	Follow uint    //关注人数
	ClassifyId uint
}