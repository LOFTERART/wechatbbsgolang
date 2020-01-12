package home

import (
	"QUZHIYOU/models"
	"QUZHIYOU/serializer"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type DiaryLikeService struct {
	DiaryId uint `json:"diary_id" form:"diary_id"`
	Type    bool `json:"type" form:"type"`
}

func (service *DiaryLikeService) LikeDiary(userid pq.Int64Array) serializer.Response {

	diary := models.Diary{
		Model: gorm.Model{
			ID: service.DiaryId,
		},
	}

	models.PG.First(&diary)

	if service.Type {
		for k, v := range diary.UserLikeId {
			if userid[0] == v {

				diary.UserLikeId = append(diary.UserLikeId[:k], diary.UserLikeId[k+1:]...)

			}
		}

		models.PG.Model(&diary).Updates(map[string]interface{}{"like": diary.Like + 1, "user_like_id": diary.UserLikeId})

	} else {
		for _, v := range diary.UserLikeId {
			if userid[0] != v {
				userid = append(userid, v)
			}

		}
		models.PG.Model(&diary).Updates(map[string]interface{}{"like": diary.Like + 1, "user_like_id": &userid})
	}

	return serializer.Response{
		Code:  0,
		Data:  nil,
		Msg:   "点赞成功",
		Error: "",
	}

}
