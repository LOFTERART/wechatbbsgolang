package models

import (
	"github.com/chenhg5/collection"
	"github.com/jinzhu/gorm"
	"strings"
)

type Comment struct {
	gorm.Model
	Name string
	Report uint  `gorm:"default: 1 "` //举报状态 1 良好  2 举报状态
	TopNum uint `gorm:"default: 0 "`//顶起来
	StepNum uint `gorm:"default: 0 "`//踩
	UserLikeId  string //点赞人id存为数组 [1,2,3,4]
	Like uint `gorm:"default: 0 "` //用户点赞数

	DiaryID uint //关联的那条日志动态
	UserID uint //哪个用户发的
	User User
}

//格式化时间戳
func (item *Comment)ForMatTime() int64 {

	return  item.CreatedAt.UnixNano() / 1e6

}



//判断用户是否点赞 根据用户id是否在这个日记中
func (info *Comment) UserIsLike(id string) bool {

	strSlice:=strings.Split(info.UserLikeId, ",")
	var array []string

	for _, v := range strSlice {
		array = append(array, v)
	}

	b := collection.Collect(array).Contains(id)

	return b

}