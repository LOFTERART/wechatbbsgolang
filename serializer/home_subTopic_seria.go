package serializer

import (
	"QUZHIYOU/models"
	"strconv"
)

type SubTopic struct {
	Id         uint   `json:"id"`
	Name       string `json:"name,omitempty"`
	SubName    string `json:"sub_name,omitempty"`
	SendNum    int   `json:"send_num,omitempty"` //发送帖子数
	Follow     int   `json:"follow,omitempty"`   //关注人数
	ClassifyId int   `json:"classify_id,omitempty"`
	BackSubName string `json:"subName,omitempty"`
	Des string `json:"des,omitempty"`
	DesPic string	`json:"desPic,omitempty"`
	Image string `json:"image,omitempty"`
}

// Build  序列化
func BuildSubTopic(item *models.SubTopic) *SubTopic {
	return &SubTopic{
		Id:         item.ID,
		Name:       item.Name,
		SubName:    item.SubName,
		SendNum:    item.SendNum,
		Follow:     item.Follow,
		ClassifyId: item.ClassifyId,
		Des: item.Des,
		DesPic: item.DesPic,
		Image:item.Image,
		BackSubName:strconv.Itoa(item.SendNum)+"条日记"+strconv.Itoa(item.Follow)+"人关注",
	}
}

// Build  序列化S
func BuildSubTopics(items []*models.SubTopic) (subTopics []*SubTopic) {
	for _, item := range items {
		act := BuildSubTopic(item)
		subTopics = append(subTopics, act)
	}
	return
}