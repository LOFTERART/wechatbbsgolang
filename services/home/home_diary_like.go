package home

import (
	"QUZHIYOU/models"
	"QUZHIYOU/serializer"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type DiaryLikeService struct {
	DiaryId uint `json:"diary_id" form:"diary_id"`
}

func (service *DiaryLikeService) LikeDiary(userid pq.Int64Array) serializer.Response {



	diary:=models.Diary{
		Model:       gorm.Model{
			ID:        service.DiaryId,
		},
	}

	models.PG.Model(&diary).Update("user_like_id",&userid)


	return serializer.Response{
		Code:  0,
		Data:  nil,
		Msg:   "点赞成功",
		Error: "",
	}

}
