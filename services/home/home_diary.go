package home

import (
	"QUZHIYOU/models"
	"QUZHIYOU/serializer"
	"fmt"
)

type ListDiaryService struct {
	Page int `form:"page" json:"page" `
	Size int `form:"size" json:"size" `
	CommunityId int `form:"communityId" json:"communityId" `
	ClassifyId int  `form:"classifyId" json:"classifyId"`
}

func (service *ListDiaryService) GetDiarys() serializer.Response {

	fmt.Println(service,"service----------")

	var diarys []*models.Diary

	total := 0

	if service.Size == 0 {
		service.Size = 10
	}
	if service.Page == 0 {
		service.Page = 1
	}

	start := (service.Page - 1) * service.Size


	//根据前端是否传递ClassifyId 返回对应的数据
	if service.ClassifyId>0{
		if err := models.PG.Where("classify_id=? AND community_id=?",service.ClassifyId,service.CommunityId).Model(models.Diary{}).Count(&total).Error; err != nil {
			return serializer.Response{
				Code:  50000,
				Msg:   "数据库连接错误",
				Error: err.Error(),
			}
		}

		if err := models.PG.Where("classify_id=? AND community_id=?",service.ClassifyId,service.CommunityId).Order("id desc").Limit(service.Size).Offset(start).Find(&diarys).Error; err != nil {
			return serializer.Response{
				Code:  50000,
				Msg:   "数据库连接错误",
				Error: err.Error(),
			}
		}

	}else {
		if err := models.PG.Where("community_id=?",service.CommunityId).Model(models.Diary{}).Count(&total).Error; err != nil {
			return serializer.Response{
				Code:  50000,
				Msg:   "数据库连接错误",
				Error: err.Error(),
			}
		}

		if err := models.PG.Where("community_id=?",service.CommunityId).Order("id desc").Limit(service.Size).Offset(start).Find(&diarys).Error; err != nil {
			return serializer.Response{
				Code:  50000,
				Msg:   "数据库连接错误",
				Error: err.Error(),
			}
		}
	}




	return serializer.BuildListResponse(serializer.BuildDiarys(diarys), uint(total))


}


