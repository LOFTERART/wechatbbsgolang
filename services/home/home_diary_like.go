package home

import (
	"QUZHIYOU/models"
	"QUZHIYOU/serializer"
	"fmt"
	"github.com/jinzhu/gorm"
	"strings"
)

type DiaryLikeService struct {
	DiaryId uint `json:"diary_id" form:"diary_id"`
	Type    bool `json:"type" form:"type"`
}

func (service *DiaryLikeService) LikeDiary(userid []string) serializer.Response {

	diary := models.Diary{
		Model: gorm.Model{
			ID: service.DiaryId,
		},
	}

	models.DB.First(&diary)
	UserLikeArrayId := strings.Split(diary.UserLikeId, ",")
	fmt.Println(UserLikeArrayId,"------UserLikeArrayId-------",diary.UserLikeId)
	if service.Type {
		for k, v := range UserLikeArrayId {
			if userid[0] == v {
				//diary.UserLikeId = append(diary.UserLikeId[:k], diary.UserLikeId[k+1:]...)
				UserLikeArrayId = append(UserLikeArrayId[:k], UserLikeArrayId[k+1:]...)
			}
		}

		str:=strings.Join(UserLikeArrayId, ";")

		diary.UserLikeId=str

		models.DB.Model(&diary).
			Updates(map[string]interface{}{"like": diary.Like - 1, "user_like_id": diary.UserLikeId})

	} else {
		for _, v := range UserLikeArrayId {
			if userid[0] != v {
				userid = append(userid, v)
			}

		}
		models.DB.Model(&diary).Updates(map[string]interface{}{"like": diary.Like + 1})
		models.DB.Exec("UPDATE diary SET user_like_id = user_like_id||?  WHERE id= ? ", userid, service.DiaryId)

	}

	return serializer.Response{
		Code:  0,
		Data:  nil,
		Msg:   "操作成功",
		Error: "",
	}

}
