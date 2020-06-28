package services

import (
	"QUZHIYOU/models"
	"QUZHIYOU/serializer"
	"github.com/jinzhu/gorm"
	"strings"
)

//定义前段传递的数据字段
type Comment struct {
	Name    string `json:"name"`
	DiaryID uint   `json:"diary_id"`
	UserID  uint   `json:"user_id"`
}

//增加Comment
func (item *Comment) CreateComment() serializer.Response {

	info := models.Comment{
		Name:    item.Name,
		DiaryID: item.DiaryID,
		UserID:  item.UserID,
	}
	models.DB.Create(&info)

	return serializer.Response{
		Code: 0,
		Data: nil,
		Msg:  "评论成功",
	}
}

//删 Comment
func (item *Comment) DelComment() serializer.Response {

	return serializer.Response{
		Code:  0,
		Data:  nil,
		Msg:   "删除成功",
		Error: "",
	}
}

//查ALL Comment
func (item *Comment) GetALLComment() serializer.Response {

	var infos []*models.Comment

	models.DB.
		Preload("User").
		Where("diary_id=?", item.DiaryID).Find(&infos)

	return serializer.Response{
		Code: 0,
		Data: serializer.BuildCommentSSerializers(infos),
		Msg:  "查询ALL成功",
	}
}

//查ONE Comment
func (item *Comment) GetONEComment() serializer.Response {

	return serializer.Response{
		Code: 0,
		Data: nil,
		Msg:  "查询ONE成功",
	}
}

//更新 Comment
func (item *Comment) UpdateComment() serializer.Response {

	return serializer.Response{
		Code: 0,
		Data: nil,
		Msg:  "更新成功",
	}
}

type LikeDairyComment struct {
	CommentID uint `json:"comment_id"`
	Type    bool `json:"type"`
}

func (service *LikeDairyComment) LikeComment(userid []string) serializer.Response {

	info := models.Comment{
		Model: gorm.Model{
			ID: service.CommentID,
		},
	}

	models.DB.First(&info)
	UserLikeArrayId := strings.Split(info.UserLikeId, ",")
	if service.Type {
		for k, v := range UserLikeArrayId {
			if userid[0] == v {
				UserLikeArrayId = append(UserLikeArrayId[:k], UserLikeArrayId[k+1:]...)
			}
		}

		str := strings.Join(UserLikeArrayId, ",")

         info.UserLikeId = str

		models.DB.Model(&info).
			Updates(map[string]interface{}{"like": info.Like - 1, "user_like_id": info.UserLikeId})

	} else {
		for _, v := range UserLikeArrayId {
			if userid[0] != v {
				userid = append(userid, v)
			}
		}
		str := strings.Join(userid, ",")

         info.UserLikeId = str

		models.DB.Model(&info).
			Updates(map[string]interface{}{"like": info.Like + 1, "user_like_id": info.UserLikeId})

	}

	return serializer.Response{
		Code:  0,
		Data:  nil,
		Msg:   "操作成功",
		Error: "",
	}

}
