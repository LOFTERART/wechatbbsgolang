package serializer

import (
	"QUZHIYOU/models"
)

//社区动态 序列化
type Diary struct {
	ID                   uint                     `json:"id"`
	Name                 string                   `json:"name"`
	IsLike               bool                     `json:"is_like"` //计算得出是否点赞
	Avatar               string                   `json:"avatar"`
	Authentication       bool                     `json:"authentication"`         //认证
	AuthenticationName   string                   `json:"authentication_name"`    //认证的称号 教师 医生
	IsShowAuthentication bool                     `json:"is_show_authentication"` //是否显示认证
	Tag                  string                   `json:"tag"`
	SubTopicInfo         *SubTopic                `json:"sub_topic_info"`
	Content              string                   `json:"content"`
	Like                 uint                     `json:"like"`
	View                 int                      `json:"views"`
	Address              string                   `json:"address"`
	Community            string                   `json:"community"` //社区名字
	Photos               string `json:"image_url_came"`
	PhotosThumb         string `json:"image_url_came_thumb"`
	CreatedAt            string                   `json:"timer"`
	SubTopicId           uint                     `json:"tagId"`
	CommunityId          uint                     `json:"communityId"`

	UserLikes string `json:"user_likes"`
	UserId    uint          `json:"user_id"`

	UserInfo *User  `json:"user_info"`   //用户信息表

	Comment []*Comment `json:"comment"`  //首页显示动态评论数
	CommentNum uint `json:"comment_num"` //首页显示动态评论数
}

func BuildDiary(item models.Diary, userId string) Diary {
	return Diary{
		ID:                   item.ID,
		Name:                 item.UserInfo.NickName,
		Authentication:       item.UserInfo.Authentication,
		AuthenticationName:   item.UserInfo.AuthenticationName,
		IsShowAuthentication: item.UserInfo.IsShowAuthentication,
		Content:              item.Content,
		Like:                 item.Like,
		IsLike:               item.UserIsLike(userId),
		View:                 item.View,
		Address:              item.Address,
		//Community:            item.CommunityInfo.Name,
		Photos:               item.Photos,
		PhotosThumb:          item.PhotosThumb,
		Tag:                  item.SubTopicInfo.Name,
		Avatar:               item.UserInfo.AvatarUrl,
		CreatedAt:            item.FormatCretaeTime(),
		SubTopicId:           item.SubTopicId,
		CommunityId:          item.CommunityId,
		UserLikes:            item.UserLikeId,
		UserId:               item.UserInfo.ID,
		SubTopicInfo:         BuildSubTopic(item.SubTopicInfo), //子主题info信息
		UserInfo:BuildUserFormat(item.UserInfo),
		CommentNum:uint(len(item.Comment)),
	}

}

func BuildDiarys(items []*models.Diary, userId string) (diarys []*Diary) {

	for _, item := range items {
		diary := BuildDiary(*item, userId)
		diarys = append(diarys, &diary)
	}

	return

}
