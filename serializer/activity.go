package serializer

import (
	"QUZHIYOU/models"
)

//活动
type TbActivity struct {
	ActivityId    int64  ` gorm:"primary_key;column:ACTIVITY_ID" json:"activityId"`   // 活动ID
	ActivityName  string ` json:"activityName,omitempty" gorm:"column:ACTIVITY_NAME"` // 活动名称
	SubName       string `json:"subName,omitempty" gorm:"column:SUB_NAME"`
	Tags          string ` json:"tags,omitempty" gorm:"column:TAGS"`                    // 首页标志：#免费 #旅游 #香山 #吃喝玩乐 #一日游
	TagsInfo      string ` json:"tagsInfo,omitempty" gorm:"column:TAGS_INFO"`           // 详情页标志：'河南同乡','免费一日游','吃喝全包','客车接送'
	Price         int64  ` json:"price" gorm:"column:PRICE"`                            // 现价
	PriceTag      string ` json:"priceTag,omitempty" gorm:"column:PRICE_TAG"`           // 首届
	OriginalPrice int64  ` json:"originalPrice,omitempty" gorm:"column:ORIGINAL_PRICE"` // 原价
	Image         string ` json:"image,omitempty" gorm:"column:IMAGE"`
	MemNum        *int64 ` json:"memNum,omitempty" gorm:"column:MEM_NUM"`                // 已报名人数
	TotalNum      int64  ` json:"totalNum,omitempty" gorm:"column:TOTAL_NUM"`            // 计划总人数
	CollectionNum int64  ` json:"collectionNum,omitempty" gorm:"column:COLLECTION_NUM"`  // 收藏数
	Status        string ` json:"status,omitempty" gorm:"column:STATUS"`                 // 活动状态：未开始/进行中/已结束
	Author        string ` json:"author,omitempty" gorm:"column:AUTHOR"`                 // 创建者
	DateAdd       string ` json:"dateAdd,omitempty" gorm:"column:DATE_ADD"`              // 活动创建时间
	DateUpdate    string ` json:"dateUpdate,omitempty" gorm:"column:DATE_UPDATE"`        // 活动更新时间
	SignStartTime string ` json:"signStartTime,omitempty" gorm:"column:SIGN_START_TIME"` // 报名开始时间
	SignEndTime   string ` json:"signEndTime,omitempty" gorm:"column:SIGN_END_TIME"`     // 报名结束时间

	ActiveStartTime string ` json:"activeStartTime,omitempty" gorm:"column:ACTIVE_START_TIME"` // 活动开始时间
	ActiveEndTime   string ` json:"activeEndTime,omitempty" gorm:"column:ACTIVE_END_TIME"`     // 活动结束时间

	ActiveTime string ` json:"active_time,omitempty" gorm:"column:ACTIVE_TIME"` // 活动时间
	Views      int64  ` json:"views,omitempty" gorm:"column:VIEWS"`             // 浏览量
	HtmlCon    string ` json:"htmlCon,omitempty" gorm:"column:HTML_CON"`

	//	关联字段
	Welfares []*models.TbWelfare `json:"welfareList,omitempty" gorm:"ForeignKey:ActivityId" `

	AddressFrom []*models.TbAddress `json:"gatherAddList,omitempty" gorm:"ForeignKey:ActivityId"`

	AddressTo []*models.TbAddress `json:"destinationList,omitempty"  gorm:"ForeignKey:ActivityId"`
}

// Build  首页序列化
func BuildActivityHome(item *models.TbActivity) *TbActivity {
	return &TbActivity{
		ActivityId:    item.ActivityId,
		ActivityName:  item.ActivityName,
		SubName:       item.SubName,
		Tags:          item.Tags,
		Price:         item.Price,
		PriceTag:      item.PriceTag,
		OriginalPrice: item.OriginalPrice,
		Image:         item.Image,
		TotalNum:      item.TotalNum,
		Status:        item.Status,
	}
}

// 首页活动结果
func BuildActivitys(items []models.TbActivity) (activitys []*TbActivity) {
	for _, item := range items {
		act := BuildActivityHome(&item)
		activitys = append(activitys, act)
	}
	return activitys
}

// Build  活动详情序列化
func BuildActivity(item *models.TbActivity) *TbActivity {
	return &TbActivity{
		ActivityId:      item.ActivityId,
		ActivityName:    item.ActivityName,
		SubName:         item.SubName,
		Tags:            item.Tags,
		TagsInfo:        item.TagsInfo,
		Price:           item.Price,
		PriceTag:        item.PriceTag,
		OriginalPrice:   item.OriginalPrice,
		Image:           item.Image,
		MemNum:          item.MemNum,
		TotalNum:        item.TotalNum,
		CollectionNum:   item.CollectionNum,
		Status:          item.Status,
		Author:          item.Author,
		SignStartTime:   item.FormatTime(item.SignStartTime, "SignStartTime"),
		SignEndTime:     item.FormatTime(item.SignEndTime, "SignEndTime"),
		ActiveStartTime: item.FormatTime(item.ActiveStartTime, "ActiveStartTime"),
		ActiveEndTime:   item.FormatTime(item.ActiveEndTime, "ActiveEndTime"),
		Views:           item.AddViewNum(),
		HtmlCon:         item.HtmlCon,
		//福利
		Welfares: item.Welfares,
		//出发地
		AddressFrom: item.FormatAddressFrom(),
		//目的地
		AddressTo: item.FormatAddressTo(),
	}
}

//活动详情
type ActivityInfoJson struct {
	ActivityInfo *TbActivity         `json:"activityInfo"`
	Banner       *[]*models.TbBanner `json:"banner"`
}

// 活动详情序列器
func ActivityInfoResponse(ActivityInfo models.TbActivity, Banner *[]*models.TbBanner) Response {
	info := BuildActivity(&ActivityInfo)
	return Response{
		Data: &ActivityInfoJson{
			ActivityInfo: info,
			Banner:       Banner,
		},
	}
}
