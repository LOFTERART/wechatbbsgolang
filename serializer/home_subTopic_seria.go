package serializer

import "QUZHIYOU/models"

type SubTopic struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	SubName    string `json:"sub_name"`
	SendNum    uint   `json:"send_num"` //发送帖子数
	Follow     uint   `json:"follow"`   //关注人数
	ClassifyId uint   `json:"classify_id"`
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