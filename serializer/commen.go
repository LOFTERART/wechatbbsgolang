package serializer

import (
	"QUZHIYOU/models"
)

type Comment struct {
	ID uint `json:"id"`
	Name string`json:"name"`
	Report uint  `json:"report"`
	TopNum uint `json:"top_num"`
	StepNum uint `json:"step_num"`
	DiaryID uint `json:"diary_id"`
	UserID uint `json:"user_id"`
	User *User `json:"user"`
}

func BuildCommentSerializer(item models.Comment) *Comment {
	return &Comment{
		ID:      item.ID,
		Name:    item.Name,
		Report:  item.Report,
		TopNum:  item.TopNum,
		StepNum: item.StepNum,
		DiaryID: item.DiaryID,
		UserID:  item.UserID,
		User:BuildUserFormat(&(item.User)),
	}

}

func BuildCommentSSerializers(item []*models.Comment) (items []*Comment) {

	for _, v := range item {
		items = append(items, BuildCommentSerializer(*v))
	}
	return

}