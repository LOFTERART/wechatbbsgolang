package models

import "github.com/jinzhu/gorm"

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

	Job string //职务信息

	Comment []Comment //评论列表

	CommunityName string //用户时所在社区信息

}
