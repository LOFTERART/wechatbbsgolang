package home

import (
	"QUZHIYOU/cache"
	"QUZHIYOU/models"
	"QUZHIYOU/serializer"
	"QUZHIYOU/serializer/code"
	"github.com/jinzhu/gorm"
	"github.com/medivhzhan/weapp/v2"
	"os"
	"time"
)

type AddDiaryService struct {
	UserId      uint
	Content     string   `form:"content" json:"content" `
	Address     string   `form:"address" json:"address" `
	Community   string   `form:"community" json:"community"`
	Photos      string `form:"photos" json:"photos"`
	PhotosThumb string `form:"photosthumb" json:"photosthumb"`
	CommunityId uint     `form:"communityId" json:"communityId" `  //社区id
	ClassifyId  uint     `form:"classifyId" json:"classifyId" `    //标签ID
	SubTopicId  uint     `form:"sub_topic_id" json:"sub_topic_id"` //标签ID
}

func (diary *AddDiaryService) AddDiary(userId uint) serializer.Response {


	token, _ := cache.RedisClient.Get(cache.WeChatAccessToken).Result()
	if len(token)<=0{

		res, err := weapp.GetAccessToken(os.Getenv("WXAPP_ID"), os.Getenv("WXSECRET"))
		if err != nil {
			// 处理一般错误信息

		}

		if err := res.GetResponseError(); err != nil {
			// 处理微信返回错误信息
		}

		// 存储微信接口凭证到redis
		cache.RedisClient.Set(cache.WeChatAccessToken, res.AccessToken, 1*time.Hour)
		token, _ = cache.RedisClient.Get(cache.WeChatAccessToken).Result()
	}


	dia := models.Diary{
		UserId:      userId,
		Content:     diary.Content,
		Address:     diary.Address,
		Photos:      diary.Photos,
		PhotosThumb: diary.Photos,
		CommunityId: diary.CommunityId,
		SubTopicId:  diary.SubTopicId,
		ClassifyId:  diary.ClassifyId,
	}

	res, _ := weapp.MSGSecCheck(token, diary.Content)


	if err := res.GetResponseError(); err !=nil {
		// 处理微信返回错误信息
		return serializer.Response{
			Code:  code.ContentRisky,
			Data:  nil,
			Msg:   res.ErrMSG,
			Error: "",
		}

	}

	//创建话题
	models.DB.Create(&dia)

	//更新tag sendNum
	 subTopic:=models.SubTopic{
		 Model: gorm.Model{
			ID: diary.SubTopicId,
		 },
	 }
	models.DB.Model(&subTopic).UpdateColumn("send_num", gorm.Expr("send_num+?", 1))

	return serializer.Response{
		Code:  0,
		Data:  nil,
		Msg:   "创建成功",
		Error: "",
	}

}
