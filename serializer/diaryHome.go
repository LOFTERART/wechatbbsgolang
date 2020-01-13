package serializer

import (
	"QUZHIYOU/models"
)

//社区动态 序列化
type Diary struct {
	ID          uint                     `json:"id"`
	Name        string                   `json:"name"`
	IsLike      bool                     `json:"is_like"` //计算得出是否点赞
	Avatar      string                   `json:"avatar"`
	Authentication bool  `json:"authentication"`//认证
	AuthenticationName string `json:"authentication_name"`//认证的称号 教师 医生
	IsShowAuthentication bool `json:"is_show_authentication"`//是否显示认证

	Content     string                   `json:"content"`
	Like        uint                     `json:"like"`
	View        int                      `json:"views"`
	CommentNum  int                      `json:"comment"`
	Address     string                   `json:"address"`
	Community   string                   `json:"community"` //社区名字
	Photos      []map[string]interface{} `json:"image_url_came"`
	PhotosThumb []map[string]interface{} `json:"image_url_came_thumb"`
	Tag         string                   `json:"tag"`
	CreatedAt   string                   `json:"timer"`
	SubTopicId  uint                     `json:"tagId"`
	CommunityId uint                     `json:"communityId"`
}

func BuildDiary(item models.Diary, userId int64) Diary {
	return Diary{
		ID:          item.ID,
		Name:        item.UserInfo.NickName,
		Authentication:item.UserInfo.Authentication,
		AuthenticationName:item.UserInfo.AuthenticationName,
		IsShowAuthentication:item.UserInfo.IsShowAuthentication,
		Content:     item.Content,
		Like:        item.Like,
		IsLike:      item.UserIsLike(userId),
		View:        item.View,
		CommentNum:  item.CommentNum,
		Address:     item.Address,
		Community:   item.CommunityInfo.Name,
		Photos:      item.FormatPhotos(item.Photos),
		PhotosThumb: item.FormatPhotos(item.PhotosThumb),
		Tag:         item.Tag,

		Avatar:      item.UserInfo.AvatarUrl,
		CreatedAt:   item.FormatCretaeTime(),
		SubTopicId:  item.SubTopicId,
		CommunityId: item.CommunityId,
	}

}

func BuildDiarys(items []*models.Diary, userId int64) (diarys []*Diary) {

	for _, item := range items {
		diary := BuildDiary(*item, userId)
		diarys = append(diarys, &diary)
	}

	return

}
