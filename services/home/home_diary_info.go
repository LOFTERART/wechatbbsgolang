package home

import (
	"QUZHIYOU/models"
	"QUZHIYOU/serializer"
)

type DiaryInfoService struct {
	Id uint `json:"id"`
}

func (service *DiaryInfoService)GetDiaryInfo() serializer.Response {

	var diaryInfo models.Diary


	models.PG.
		Preload("UserInfo").
		First(&diaryInfo)

	return serializer.Response{
		Code:  0,
		Data:  &diaryInfo,
		Msg:   "",
		Error: "",
	}
}