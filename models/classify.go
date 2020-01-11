package models

type Classify struct {
	ID     int64 `gorm:"primary_key"`
	Name   string
	Image  string
	Type   string
	Des    string
	DesPic string
	SubTopics []*SubTopic ` gorm:"ForeignKey:ClassifyId" `
}
