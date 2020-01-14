package models

import (
	"github.com/chenhg5/collection"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"os"
)

//社区动态

type Diary struct {
	gorm.Model
	Content     string
	Like        uint //用户点赞数
	View        int
	CommentNum  int
	Address     string
	Photos      pq.StringArray `gorm:"type:varchar(300)[]"`
	PhotosThumb pq.StringArray `gorm:"type:varchar(300)[]"`
	Tag         string
	ClassifyId  uint          //属于哪个大标签
	UserLikeId  pq.Int64Array `gorm:"type:varchar(300)[]"` //点赞人id存为数组

	//外键信息
	UserInfo *User ` gorm:"ForeignKey:UserId" `
	UserId   uint

	CommunityInfo *Communitys ` gorm:"ForeignKey:CommunityId" `
	CommunityId   uint

	SubTopicInfo *SubTopic ` gorm:"ForeignKey:SubTopicId" `
	SubTopicId   uint
}

var timeLayoutStr = "2006/01/02 15:04"

//格式化create
func (Diary *Diary) FormatCretaeTime() string {
	ts := Diary.CreatedAt.Format(timeLayoutStr) //time转string
	return ts
}

//格式化photos
func (Diary *Diary) FormatPhotos(photo []string) (photos []map[string]interface{}) {

	for _, v := range photo {
		maPhoto := make(map[string]interface{})
		maPhoto["url"] = os.Getenv("PIC_TOKEN") + v
		photos = append(photos, maPhoto)
	}

	return

}

//判断用户是否点赞 根据用户id是否在这个日记中
func (Diary *Diary) UserIsLike(id int64) bool {

	var array []int64

	for _, v := range Diary.UserLikeId {
		array = append(array, v)
	}

	b := collection.Collect(array).Contains(id)

	return b

}
