package models

//活动福利
type TbWelfare struct {
	WelfareID  int64  `gorm:"primary_key;column:WELFARE_ID" json:"welfareId"` // 福利ID
	Tag        string `gorm:"column:TAG" json:"tag"`                          // 福利类型
	Des        string ` gorm:"column:DES"json:"des"`                          // 福利详细描述
	ActivityId int64  `gorm:"index;column:ACTIVITY_ID" json:"-"`              // 活动ID

}
