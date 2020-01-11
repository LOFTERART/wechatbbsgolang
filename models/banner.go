package models

//活动图片
type TbBanner struct {
	BannerID   int64  ` gorm:"primary_key;column:BANNER_ID" json:"-"`
	ActivityId int64  ` gorm:"column:ACTIVITY_ID" json:"-"` // 活动ID
	URL        string `  gorm:"column:URL" json:"url"`      // 图片URL
}
