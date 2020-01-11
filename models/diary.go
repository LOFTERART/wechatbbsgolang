package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"os"
)

//社区动态

type Diary struct {
	gorm.Model
	Avatar string
	Name       string
	Content    string
	Like       int
	IsLike     bool
	View       int
	Auth       string
	CommentNum int
	Address    string
	Community  string
	Photos      pq.StringArray `gorm:"type:varchar(300)[]"`
	PhotosThumb      pq.StringArray `gorm:"type:varchar(300)[]"`
	Tag     string
	Status string
	Specialist bool
	CommunityId uint
	SubTopicId uint //属于哪个标签
	ClassifyId uint //属于哪个大标签
}

var timeLayoutStr = "2006/01/02 15:04"


//格式化create
func (Diary *Diary) FormatCretaeTime() string {
	ts := Diary.CreatedAt.Format(timeLayoutStr) //time转string
	return ts
}

func (Diary *Diary) FormatPhotos(photo []string) (photos []map[string]interface{}) {

	for _,v:=range photo{
		maPhoto:=make(map[string]interface{})
		maPhoto["url"]=os.Getenv("PIC_TOKEN")+v
		photos=append(photos,maPhoto)
	}

	return

}