package home

import (
	"QUZHIYOU/models"
	"QUZHIYOU/serializer"
)

type DiaryInfoService struct {
	Id int `form:"id" json:"id" `
}

func (service *DiaryInfoService) GetDiaryInfo(userId uint) serializer.Response {

	var DiaryInfo models.Diary

	models.PG.
		Where("id=?", uint(service.Id)).
		First(&DiaryInfo)

	var users []models.User

	var ids []uint

	for _, v := range DiaryInfo.UserLikeId {
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
