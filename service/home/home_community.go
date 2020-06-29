package home

import (
	"QUZHIYOU/models"
	"QUZHIYOU/serializer"
)

type HomeCommunityService struct {
}

func (service *HomeCommunityService) GetCommunity() serializer.Response {

	var communitys []models.Communitys

	models.DB.Find(&communitys)

	l1 := serializer.BuildCommunitys(communitys)
	return serializer.Response{
		Code:  0,
		Data:  &l1,
		Msg:   "",
		Error: "",
	}

}
