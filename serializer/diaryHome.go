package serializer

import (
	"QUZHIYOU/models"
)

//社区动态 序列化
type Diary struct {
	ID          uint           `json:"id"`
	Name        string         `json:"name"`
	Content     string         `json:"content"`
	Like        int            `json:"like"`
	IsLike      bool           `json:"is_like"`
	View        int            `json:"views"`
	Auth        string         `json:"Auth"`
	CommentNum  int            `json:"comment"`
	Address     string         `json:"address"`
	Community   string         `json:"community"`
	Photos      []map[string]interface{} `json:"image_url_came"`
	PhotosThumb []map[string]interface{} `json:"image_url_came_thumb"`
	Tag         string         `json:"tag"`
	Status string `json:"status"`
	Specialist bool `json:"specialist"`
	Avatar string `json:"avatar"`
	CreatedAt   string `json:"timer"`
	SubTopicId  uint `json:"tagId"`
	CommunityId uint `json:"communityId"`
}

func BuildDiary(item models.Diary) Diary {
	return Diary{
		ID:          item.ID,
		Name:        item.Name,
		Content:     item.Content,
		Like:        item.Like,
		IsLike:      item.UserIsLike(int64(18)),
		View:        item.View,
		Auth:        item.Auth,
		CommentNum:  item.CommentNum,
		Address:     item.Address,
		Community:   item.Community,
		Photos:      item.FormatPhotos(item.Photos),
		PhotosThumb: item.FormatPhotos(item.PhotosThumb),
		Tag:         item.Tag,
		Status:         item.Status,
		Specialist:item.Specialist,
		Avatar:item.Avatar,
		CreatedAt:item.FormatCretaeTime(),
		SubTopicId:item.SubTopicId,
		CommunityId :item.CommunityId,
	}

}

func BuildDiarys(items []*models.Diary) (diarys []*Diary) {

	for _, item := range items {
		diary := BuildDiary(*item)
		diarys = append(diarys, &diary)
	}

	return

}
