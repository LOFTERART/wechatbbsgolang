package home

import (
	"QUZHIYOU/models"
	"QUZHIYOU/serializer"
)

type SubTopicService struct {

}

func (ad *SubTopicService)GetSubTopic() serializer.Response {

	var SubTopics []*models.SubTopic

	models.PG.Find(&SubTopics)

	return serializer.Response{
		Code:  0,
		Data:  serializer.BuildSubTopics(SubTopics),
		Msg:   "",
		Error: "",
	}




}