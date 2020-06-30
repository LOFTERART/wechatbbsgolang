package serializer

import (
	"QUZHIYOU/models"
)

type Comment struct {
	ID uint `json:"id"`
	Name string`json:"name,omitempty"`
	Report uint  `json:"report,omitempty"`
	TopNum uint `json:"top_num,omitempty"`
	StepNum uint `json:"step_num,omitempty"`
	DiaryID uint `json:"diary_id,omitempty"`
	UserID uint `json:"user_id,omitempty"`
	User *User `json:"user,omitempty"`
	CreateAt int64 `json:"create_at,omitempty"`
	Like uint `json:"like"`  //点赞数
	IsLike   bool    `json:"is_like,omitempty"` //计算得出是否点赞
}

func BuildCommentSerializer(item models.Comment,userID string) *Comment {
	return &Comment{
		ID:      item.ID,
		Name:    item.Name,
		Report:  item.Report,
		TopNum:  item.TopNum,
		StepNum: item.StepNum,
		DiaryID: item.DiaryID,
		UserID:  item.UserID,
		User:BuildUserFormat(&(item.User)),
		CreateAt: item.ForMatTime(),
		Like: item.Like,
		IsLike:item.UserIsLike(userID),  //判断当前用户是否对评论点赞

	}

}

func BuildCommentSSerializers(item []*models.Comment,userID string) (items []*Comment) {

	for _, v := range item {
		items = append(items, BuildCommentSerializer(*v,userID))
	}
	return

}