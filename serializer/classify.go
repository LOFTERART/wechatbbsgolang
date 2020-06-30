package serializer

import "QUZHIYOU/models"

type Classify struct {
	ID      int64       `json:"id"`
	Name    string      `json:"name",omitempty`
	Image   string      `json:"image",omitempty`
	Type    string      `json:"type",omitempty`
	Des     string      `json:"des,omitempty"`
	DesPic  string      `json:"desPic,omitempty"`
	SubTops []*SubTopic `json:"subTops"`
}

//单行序列化
func BuildClassify(item *models.Classify) *Classify {

	return &Classify{
		ID:      item.ID,
		Name:    item.Name,
		Image:   item.Image,
		Type:    item.Type,
		Des:     item.Des,
		DesPic:  item.DesPic,
		SubTops: BuildSubTopics(item.SubTopics),
	}
}

//多行序列化
func BuildClassifys(item []*models.Classify) (classify []*Classify) {

	for _, v := range item {
		classify = append(classify, BuildClassify(v))
	}
	return

}
