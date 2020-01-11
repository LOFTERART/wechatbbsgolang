package serializer

import "QUZHIYOU/models"

type Classify struct {
	ID   int64 `json:"id"`
	Name string `json:"name"`
	Image string `json:"image"`
	Type string `json:"type"`
	Des string `json:"des"`
	DesPic string `json:"desPic"`
	SubTops []*SubTopic `json:"sub_tops"`
}

//单行序列化
func BuildClassify(item *models.Classify) *Classify {

	return &Classify{
		ID:   item.ID,
		Name: item.Name,
		Image: item.Image,
		Type:item.Type,
		Des:item.Des,
		DesPic:item.DesPic,
		SubTops:BuildSubTopics(item.SubTopics),
	}
}

//多行序列化
func BuildClassifys(item []*models.Classify) (classify []*Classify) {

	for _, v := range item {
		classify = append(classify, BuildClassify(v))
	}
	return

}
