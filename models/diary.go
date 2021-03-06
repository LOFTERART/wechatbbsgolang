package models

import (
	"github.com/chenhg5/collection"
	"github.com/disintegration/imaging"
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"image"
	"image/color"
	"log"
	"os"
	"strings"
)

//社区动态

type Diary struct {
	gorm.Model
	Content     string
	Like        uint //用户点赞数
	View        int
	Address     string
	Photos      string
	PhotosThumb string
	Tag         string
	ClassifyId  uint   //属于哪个大标签
	UserLikeId  string //点赞人id存为数组

	//外键信息
	UserInfo *User ` gorm:"ForeignKey:UserId" `
	UserId   uint

	CommunityId uint //社区ID

	SubTopicInfo *SubTopic ` gorm:"ForeignKey:SubTopicId" `
	SubTopicId   uint

	Comment []*Comment
}

var timeLayoutStr = "2006/01/02 15:04"

//格式化create
func (Diary *Diary) FormatCretaeTime() string {
	ts := Diary.CreatedAt.Format(timeLayoutStr) //time转string
	return ts
}

//格式化photos
func (Info *Diary) FormatPhotos() (photos []string) {

	if len(Info.Photos) <= 0 {
		photos = []string{}
		return
	} else {
		photoArray := strings.Split(Info.Photos, "￥")

		for _, v := range photoArray {
			photos = append(photos, os.Getenv("IMGADDRESS")+v)
		}
		return
	}

}

//格式化photosthumb
func (Info *Diary) FormatPhotosThumb() (photos []string) {

	if len(Info.PhotosThumb) <= 0 {
		photos = []string{}
		return
	} else {

		dir, _ := os.Getwd()
		photoArray := strings.Split(Info.PhotosThumb, "￥")

		for _, v := range photoArray {
			src, err := imaging.Open(dir + "/static/" + v)
			if err != nil {
				log.Fatalf("failed to open image: %v", err)
			}
			src = imaging.Resize(src, 800, 800, imaging.NearestNeighbor)
			dst := imaging.New(400, 400, color.NRGBA{255, 255, 255, 0})
			dst = imaging.Paste(dst, src, image.Pt(0, 0))
			imaging.Save(dst, dir+"/static/thumb/"+v)
			photos = append(photos, os.Getenv("IMGADDRESSTHUMB")+v)
		}
		return

	}

}

//判断用户是否点赞 根据用户id是否在这个日记中
func (Diary *Diary) UserIsLike(id string) bool {

	strSlice := strings.Split(Diary.UserLikeId, ",")
	var array []string

	for _, v := range strSlice {
		array = append(array, v)
	}

	b := collection.Collect(array).Contains(id)

	return b

}

type DiaryBACK struct {
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
