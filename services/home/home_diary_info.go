package home

import (
	"QUZHIYOU/models"
	"QUZHIYOU/serializer"
	"strconv"
	"strings"
)

type DiaryInfoService struct {
	Id int `form:"id" json:"id" `
}

func (service *DiaryInfoService) GetDiaryInfo(userId uint) serializer.Response {

	var DiaryInfo models.Diary

	models.DB.Where("id=?", uint(service.Id)).First(&DiaryInfo)


	//用户信息  用来装在用户的图片
	var users []*models.User

	UserLikeArrayId := strings.Split(DiaryInfo.UserLikeId, ",")
	//设置已经点赞的用户id
	var ids []uint
	for _, v := range UserLikeArrayId {
		v1,_:=strconv.Atoi(v)
		ids = append(ids, uint(v1))
	}

	models.DB.Where("id in (?)", ids).Find(&users)

	//res:=serializer.BuildHomeDiaryInfoUserPics(users)
	res:=serializer.BuildUserSFormat(users)


	return serializer.Response{
		Code:  0,
		Data:  map[string]interface{}{"total":len(ids),"items":&res},
		Msg:   "",
		Error: "",
	}
}
