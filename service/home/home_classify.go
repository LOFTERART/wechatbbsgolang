package home

import (
	"QUZHIYOU/models"
	"QUZHIYOU/serializer"
)

type HomeClassify struct {
}

func (classify *HomeClassify) GetClassify() serializer.Response {

	var classifys []*models.Classify

	models.DB.Preload("SubTopics").Limit(10).Find(&classifys)

	return serializer.Response{
		Data:  serializer.BuildClassifys(classifys),
	}

}
