package models

import "github.com/jinzhu/gorm"

//用户
type User struct {
	gorm.Model
	OpenId      string
	SessionKey  string
	UnionId     string
	UserName    string
	UserFrom    string
	AccessToken string
	Auth        string //职业认证
	Specialist  bool   //是否认证
	NickName string
	AvatarUrl string
	Province string
	City string
	Country string
	Gender int

}
