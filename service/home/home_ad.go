package home

import (
	"QUZHIYOU/models"
	"QUZHIYOU/serializer"
)

type AdService struct {
}

func (ad *AdService) GetAds() serializer.Response {

	var ads []*models.Ad

	models.DB.Find(&ads)

	return serializer.Response{
		Code:  0,
		Data:  serializer.BuildAdSerializers(ads),
		Msg:   "",
		Error: "",
	}

}
