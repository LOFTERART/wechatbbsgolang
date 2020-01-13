package serializer

import (
	"QUZHIYOU/models"
	"strconv"
)

type SubTopic struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	SubName    string `json:"sub_name"`
	SendNum    int   `json:"send_num"` //发送帖子数
	Follow     int   `json:"follow"`   //关注人数
	ClassifyId int   `json:"classify_id"`
	BackSubName string `json:"subName"`
	Desc string `json:"des"`
	DescPic string	`json:"desPic"`
	Image string `json:"image"`
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
		Desc: item.Desc,
		DescPic: item.DescPic,
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