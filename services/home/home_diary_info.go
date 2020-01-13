package home

import (
	"QUZHIYOU/models"
	"QUZHIYOU/serializer"
	"fmt"
)

type DiaryInfoService struct {
	Id int `json:"id"`
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


	return serializer.Response{
		Code:  0,
		Data:  serializer.BuildDiary(DiaryInfo,int64(userId)),
		Msg:   "",
		Error: "",
	}
}