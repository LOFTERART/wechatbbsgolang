package home

import (
	"QUZHIYOU/models"
	"QUZHIYOU/serializer"
)

type ListDiaryService struct {
	Page        int `form:"page" json:"page" `
	Size        int `form:"size" json:"size" `
	CommunityId int `form:"communityId" json:"communityId" `
	ClassifyId  int `form:"classifyId" json:"classifyId"`
	SubTopicId  int `form:"sub_topic_id" json:"sub_topic_id"`
	UserId      int `form:"user_id" json:"user_id"` //传递userid 进入用户胡中心
}

func (service *ListDiaryService) GetDiarys(userId string) serializer.Response {

	var diarys []*models.Diary

	total := 0

	if service.Size == 0 {
		service.Size = 10
	}
	if service.Page == 0 {
		service.Page = 1
	}

	start := (service.Page - 1) * service.Size

	//转换传递用户ID  str=>int
	//userIDint,_:=strconv.Atoi(userId)

	//根据前端是否传递ClassifyId 返回对应的数据
	if service.ClassifyId > 0 {
		if err := models.DB.Where("classify_id=? AND community_id=?", service.ClassifyId, service.CommunityId).
			Model(models.Diary{}).Count(&total).Error; err != nil {
			return serializer.Response{
				Code:  50000,
				Msg:   "数据库连接错误",
				Error: err.Error(),
			}
		}

		if err := models.DB.
			Preload("UserInfo").
			Preload("Comment").
			Preload("SubTopicInfo").
			Where("classify_id=? AND community_id=?", service.ClassifyId, service.CommunityId).
			Order("id desc").Limit(service.Size).Offset(start).Find(&diarys).Error; err != nil {
			return serializer.Response{
				Code:  50000,
				Msg:   "数据库连接错误",
				Error: err.Error(),
			}
		}

	} else if service.SubTopicId > 0 {
		if err := models.DB.Where(" community_id=? AND sub_topic_id=?", service.CommunityId, service.SubTopicId).
			Model(models.Diary{}).Count(&total).Error; err != nil {
			return serializer.Response{
				Code:  50000,
				Msg:   "数据库连接错误",
				Error: err.Error(),
			}
		}

		if err := models.DB.
			Preload("UserInfo").
			Preload("Comment").
			Preload("SubTopicInfo").
			Where(" community_id=? AND sub_topic_id=?", service.CommunityId, service.SubTopicId).
			Order("id desc").Limit(service.Size).Offset(start).Find(&diarys).Error; err != nil {
			return serializer.Response{
				Code:  50000,
				Msg:   "数据库连接错误",
				Error: err.Error(),
			}
		}

	} else if service.UserId > 0 {
		if err := models.DB.Where(" user_id=? ", service.UserId).Model(models.Diary{}).
			Count(&total).Error; err != nil {
			return serializer.Response{
				Code:  50000,
				Msg:   "数据库连接错误",
				Error: err.Error(),
			}
		}
		//Where(" user_id=? AND community_id=?", service.UserId, service.CommunityId).
		if err := models.DB.
			Where(" user_id=?", service.UserId).
			Preload("UserInfo").
			Preload("Comment").
			Preload("SubTopicInfo").
			Order("id desc").Limit(service.Size).Offset(start).Find(&diarys).Error; err != nil {
			return serializer.Response{
				Code:  50000,
				Msg:   "数据库连接错误",
				Error: err.Error(),
			}
		}


	} else {

		if service.CommunityId==99999{

			if err := models.DB.Model(models.Diary{}).
				Count(&total).Error; err != nil {
				return serializer.Response{
					Code:  50000,
					Msg:   "数据库连接错误",
					Error: err.Error(),
				}
			}

			if err := models.DB.
				Preload("UserInfo").
				Preload("Comment").
				Preload("SubTopicInfo").
				Order("id desc").
				Limit(service.Size).Offset(start).Find(&diarys).Error; err != nil {
				return serializer.Response{
					Code:  50000,
					Msg:   "数据库连接错误",
					Error: err.Error(),
				}
			}
		}else {
			if err := models.DB.Where("community_id=?", service.CommunityId).Model(models.Diary{}).
				Count(&total).Error; err != nil {
				return serializer.Response{
					Code:  50000,
					Msg:   "数据库连接错误",
					Error: err.Error(),
				}
			}

			if err := models.DB.
				Preload("UserInfo").
				Preload("Comment").
				Preload("SubTopicInfo").
				Where("community_id=?", service.CommunityId).
				Order("id desc").
				Limit(service.Size).Offset(start).Find(&diarys).Error; err != nil {
				return serializer.Response{
					Code:  50000,
					Msg:   "数据库连接错误",
					Error: err.Error(),
				}
			}
		}


	}

	return serializer.BuildListResponse(serializer.BuildDiarys(diarys, userId), uint(total))

}
