package home

import (
	"QUZHIYOU/models"
	"QUZHIYOU/serializer"
)

type SubTopicService struct {
	ClassifyId uint `form:"classifyId" json:"classifyId"`
}

func (service *SubTopicService) GetSubTopic() serializer.Response {

	var SubTopics []*models.SubTopic

	models.DB.Where("classify_id=?", service.ClassifyId).Find(&SubTopics)

	return serializer.Response{
		Code:  0,
		Data:  serializer.BuildSubTopics(SubTopics),
		Msg:   "",
		Error: "",
	}

}
