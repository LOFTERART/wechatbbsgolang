package models

//活动地址
type TbAddress struct {
	AddressId   int64   `gorm:"primary_key;column:ADDRESS_ID" json:"addressId"` // 地址ID
	AddressName string  `gorm:"column:ADDRESS_NAME" json:"addressName"`         // 地址
	Address     string  `gorm:"column:ADDRESS" json:"address"`                  // 地址详情
	Type        int64   `gorm:"column:TYPE" json:"type"`                        // 地址类型：1-集合地/2-目的地
	Lat         float64 `gorm:"column:LAT" json:"lat"`                          // 纬度
	Lng         float64 `gorm:"column:LNG" json:"lng"`                          // 经度
	ActivityId  int64   `gorm:"INDEX;column:ACTIVITY_ID" json:"-"`              // 活动ID
}
