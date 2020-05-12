package activity

import (
	"QUZHIYOU/models"
	"QUZHIYOU/serializer"
	"QUZHIYOU/utils"
	"github.com/jinzhu/gorm"
)

//活动详情
type ActivityInfo struct {
	ActivityId string `form:"activityId" json:"activityId"`
}

func (activityInfo *ActivityInfo) GetActivityInfo() serializer.Response {

	actiInfo := models.TbActivity{
		ActivityId: utils.String2Int64(activityInfo.ActivityId),
	}

	var banners []*models.TbBanner

	models.DB.Find(&banners, "ACTIVITY_ID=?", utils.String2Int64(activityInfo.ActivityId))

	err := models.DB.Debug().
		Preload("Welfares").
		Preload("AddressFrom").
		Preload("AddressTo").
		First(&actiInfo).
		UpdateColumn("VIEWS", gorm.Expr("VIEWS + ?", 1)).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "活动不存在",
			Error: err.Error(),
		}
	}

	return serializer.ActivityInfoResponse(actiInfo, &banners)

}
