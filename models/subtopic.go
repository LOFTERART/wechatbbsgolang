package models

import "github.com/jinzhu/gorm"

type SubTopic struct {
	gorm.Model
	Name       string
	SubName    string
	SendNum    int `gorm:"default: 0 "` //发送帖子数
	Follow     int `gorm:"default: 0 "` //关注人数
	ClassifyId int
	//Desc string
	//DescPic string  //背景大图
	Des    string
	DesPic string
	Image  string //主题icon
}
