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
	

	if service.Size == 0 {
		service.Size = 10
	}
	if service.Page == 0 {
		service.Page = 1
	}

	start := (service.Page - 1) * service.Size

	//根据前端是否传递king 分类ID  ClassifyId 返回对应的数据
	if service.ClassifyId > 0 {

		if service.CommunityId == 99999 {
			models.DB.
				Preload("UserInfo").
				Preload("Comment").
				Preload("SubTopicInfo").
				Where("classify_id=? ", service.ClassifyId).
				Order("id desc").Limit(service.Size).Offset(start).Find(&diarys)
		} else {
			models.DB.
				Preload("UserInfo").
				Preload("Comment").
				Preload("SubTopicInfo").
				Where("classify_id=? AND community_id=?", service.ClassifyId, service.CommunityId).
				Order("id desc").Limit(service.Size).Offset(start).Find(&diarys)
		}

	} else if service.SubTopicId > 0 {
		//根据前端是否传递king下面对应的子主题 返回对应的数据
		    models.DB.
			Preload("UserInfo").
			Preload("Comment").
			Preload("SubTopicInfo").
			Where(" community_id=? AND sub_topic_id=?", service.CommunityId, service.SubTopicId).
			Order("id desc").
			Limit(service.Size).
			Offset(start).
			Find(&diarys)

	} else if service.UserId > 0 {
       //个人中心的对应动态
		models.DB.
			Where(" user_id=?", service.UserId).
			Preload("UserInfo").
			Preload("Comment").
			Preload("SubTopicInfo").
			Order("id desc").Limit(service.Size).Offset(start).Find(&diarys)

	} else {
        //寻找社区对应的动态
		if service.CommunityId == 99999 {

			models.DB.
				Preload("UserInfo").
				Preload("Comment").
				Preload("SubTopicInfo").
				Order("id desc").
				Limit(service.Size).Offset(start).Find(&diarys)

		} else {

			    models.DB.
				Preload("UserInfo").
				Preload("Comment").
				Preload("SubTopicInfo").
				Where("community_id=?", service.CommunityId).
				Order("id desc").
				Limit(service.Size).Offset(start).Find(&diarys)
		}

	}

	return serializer.Response{
		Code:  0,
		Data:  serializer.BuildDiarys(diarys, userId),
	}

}
