package serializer

import (
	"QUZHIYOU/models"
)

//社区动态 序列化
type Diary struct {
	ID                   uint                     `json:"id"`
	Name                 string                   `json:"name,omitempty"`
	IsLike               bool                     `json:"is_like,omitempty"` //计算得出是否点赞
	Avatar               string                   `json:"avatar,omitempty"`
	Authentication       bool                     `json:"authentication,omitempty"`         //认证
	AuthenticationName   string                   `json:"authentication_name,omitempty"`    //认证的称号 教师 医生
	IsShowAuthentication bool                     `json:"is_show_authentication,omitempty"` //是否显示认证
	Tag                  string                   `json:"tag,omitempty"`
	SubTopicInfo         *SubTopic                `json:"sub_topic_info,omitempty"`
	Content              string                   `json:"content,omitempty"`
	Like                 uint                     `json:"like"`
	View                 int                      `json:"views,omitempty"`
	Address              string                   `json:"address,omitempty"`
	Photos               []string `json:"image_url_came,omitempty"`
	PhotosThumb          []string `json:"image_url_came_thumb,omitempty"`
	CreatedAt            string                   `json:"timer,omitempty"`
	SubTopicId           uint                     `json:"tagId,omitempty"`
	CommunityId          uint                     `json:"communityId,omitempty"`

	UserLikes string `json:"user_likes,omitempty"`
	UserId    uint          `json:"user_id,omitempty"`

	UserInfo *User  `json:"user_info,omitempty"`   //用户信息表

	Comment []*Comment `json:"comment,omitempty"`  //首页显示动态评论数
	CommentNum uint `json:"comment_num,omitempty"` //首页显示动态评论数

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
		Photos:               item.FormatPhotos(),
		PhotosThumb:          item.FormatPhotosThumb(),
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
