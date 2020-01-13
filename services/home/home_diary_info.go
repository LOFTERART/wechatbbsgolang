package home

import (
	"QUZHIYOU/models"
	"QUZHIYOU/serializer"
	"fmt"
)

type DiaryInfoService struct {
	Id int `form:"id" json:"id" `
}

func (service *DiaryInfoService)GetDiaryInfo(userId uint) serializer.Response {

	fmt.Println(service.Id,"--------sweeee")

	var DiaryInfo models.Diary


	models.PG.
		Where("id=?",uint(89)).
		Preload("UserInfo").
		Preload("SubTopicInfo").
		Preload("CommunityInfo").
		First(&DiaryInfo)


	var users []models.User

	var ids  []uint

	for _,v:=range DiaryInfo.UserLikeId{
		ids = append(ids, uint(v))
	}

	models.PG.Where("id in (?)", ids).Find(&users)



	return serializer.Response{
		Code:  0,
		Data:  &users,
		Msg:   "",
		Error: "",
	}
}