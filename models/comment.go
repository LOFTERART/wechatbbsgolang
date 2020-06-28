package models

import "github.com/jinzhu/gorm"

type Comment struct {
	gorm.Model
	Name string
	Report uint  `gorm:"default: 1 "` //举报状态 1 良好  2 举报状态
	TopNum uint `gorm:"default: 0 "`//顶起来
	StepNum uint `gorm:"default: 0 "`//踩
	DiaryID uint //关联的那条日志动态
	UserID uint //哪个用户发的
	User User
}
