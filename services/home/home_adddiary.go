package home

import (
	"QUZHIYOU/models"
	"QUZHIYOU/serializer"
	"github.com/jinzhu/gorm"
)

type AddDiaryService struct {
	Avatar    string `form:"avatar" json:"avatar" `
	Name      string `form:"name" json:"name" `
	Content   string `form:"content" json:"content" `
	Address   string `form:"address" json:"address" `
	Community string `form:"community" json:"community"`
	Photos []string `form:"photos" json:"photos"`
	Tag string `form:"tag" json:"tag"`
	PhotosThumb []string `form:"photosthumb" json:"photosthumb"`
	CommunityId uint `form:"communityId" json:"communityId" `  //社区id
	ClassifyId uint `form:"classifyId" json:"classifyId" `    //标签ID
	SubTopicId uint `form:"sub_topic_id" json:"sub_topic_id"`    //标签ID
}

func (diary *AddDiaryService) AddDiary() serializer.Response {

	dia := models.Diary{
		Content:   diary.Content,
		Address:   diary.Address,
		Community: diary.Community,
		Photos:diary.Photos,
		PhotosThumb:diary.Photos,
		Tag:diary.Tag,
		CommunityId:diary.CommunityId,
		SubTopicId:diary.SubTopicId,
		ClassifyId:diary.ClassifyId,
	}

	//创建话题
	models.PG.Create(&dia)
	//更新tag sendNum
	 var subTopic models.SubTopic
	subTopic.ID=diary.ClassifyId
	models.PG.Model(&subTopic).UpdateColumn("send_num",gorm.Expr("send_num + ?", 1))

	return serializer.Response{
		Code:  0,
		Data:  nil,
		Msg:   "创建成功",
		Error: "",
	}

}
