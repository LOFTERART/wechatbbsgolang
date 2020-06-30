package models

import (
	"github.com/chenhg5/collection"
	"github.com/jinzhu/gorm"
	"strings"
)

//用户
type User struct {
	gorm.Model
	OpenId               string
	SessionKey           string
	UnionId              string
	UserName             string
	UserFrom             string
	AccessToken          string
	Auth                 string //职业认证
	Specialist           bool   //是否认证
	NickName             string
	AvatarUrl            string
	Province             string
	City                 string
	Country              string
	Gender               int
	Authentication       bool   //认证
	AuthenticationName   string //认证的称号 教师 医生
	IsShowAuthentication bool   //是否显示认证
	Follow string   //我都被谁给关注了
	MyFollow string  //我都关注了谁


	Job string //职务信息

	Comment []Comment //评论列表

	CommunityName string //用户时所在社区信息

}


//判断用户是否点赞 根据用户id是否在这个日记中
func (info *User) UserIsLike(id string) bool {

	strSlice := strings.Split(info.Follow, ",")
	var array []string

	for _, v := range strSlice {
		array = append(array, v)
	}

	b := collection.Collect(array).Contains(id)

	return b

}