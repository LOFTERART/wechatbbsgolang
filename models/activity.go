package models

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
	MemNum        *int64 ` json:"memNum" gorm:"column:MEM_NUM"`                         // 已报名人数
	TotalNum      int64  ` json:"totalNum,omitempty" gorm:"column:TOTAL_NUM"`           // 计划总人数
	CollectionNum int64  ` json:"collectionNum,omitempty" gorm:"column:COLLECTION_NUM"` // 收藏数
	Status        string ` json:"status,omitempty" gorm:"column:STATUS"`                // 活动状态：未开始/进行中/已结束
	Author        string ` json:"author,omitempty" gorm:"column:AUTHOR"`                // 创建者
	DateAdd       string ` json:"dateAdd,omitempty" gorm:"column:DATE_ADD"`             // 活动创建时间
	DateUpdate    string ` json:"dateUpdate,omitempty" gorm:"column:DATE_UPDATE"`       // 活动更新时间

	SignStartTime string ` json:"signStartTime,omitempty" gorm:"column:SIGN_START_TIME"` // 报名开始时间
	SignEndTime   string ` json:"signEndTime,omitempty" gorm:"column:SIGN_END_TIME"`     // 报名结束时间

	ActiveStartTime string ` json:"activeStartTime,omitempty" gorm:"column:ACTIVE_START_TIME"` // 活动开始时间
	ActiveEndTime   string ` json:"activeEndTime,omitempty" gorm:"column:ACTIVE_END_TIME"`     // 活动结束时间

	ActiveTime string ` json:"active_time,omitempty" gorm:"column:ACTIVE_TIME"` // 活动时间
	Views      int64  ` json:"views,omitempty" gorm:"column:VIEWS"`             // 浏览量
	HtmlCon    string ` json:"htmlCon,omitempty" gorm:"column:HTML_CON"`

	//	关联字段
	Welfares []*TbWelfare `json:"welfareList,omitempty" gorm:"ForeignKey:ActivityId" `

	AddressFrom []*TbAddress `json:"gatherAddList,omitempty" gorm:"ForeignKey:ActivityId"`

	AddressTo []*TbAddress `json:"destinationList,omitempty"  gorm:"ForeignKey:ActivityId"`
}

//数据库先查后更新view 界面更新views+1

func (TbActivity *TbActivity) AddViewNum() int64 {
	return TbActivity.Views + 1
}

//活动出发地址格式化
func (TbActivity *TbActivity) FormatAddressFrom() []*TbAddress {

	var addFrom []*TbAddress
	for _, v := range TbActivity.AddressFrom {
		if v.Type == 1 {
			addFrom = append(addFrom, v)
		}
	}
	return addFrom

}

//活动目的地址格式化
func (TbActivity *TbActivity) FormatAddressTo() []*TbAddress {
	var addTo []*TbAddress
	for _, v := range TbActivity.AddressTo {
		if v.Type == 2 {
			addTo = append(addTo, v)
		}
	}
	return addTo

}

//活动日期格式化
func (TbActivity *TbActivity) FormatTime(time string, typeTime string) string {

	if typeTime == "SignStartTime" {
		TbActivity.SignStartTime = time[:10]
		return TbActivity.SignStartTime
	} else if typeTime == "SignEndTime" {
		TbActivity.SignEndTime = time[:10]
		return TbActivity.SignEndTime
	} else if typeTime == "ActiveStartTime" {
		TbActivity.ActiveStartTime = time[:10]
		return TbActivity.ActiveStartTime
	} else if typeTime == "ActiveEndTime" {
		TbActivity.ActiveEndTime = time[:10]
		return TbActivity.ActiveEndTime
	} else {
		return ""
	}

}

//解耦格式化目的地
func (TbActivity *TbActivity) FormatAddress1(active *TbActivity) ([]*TbAddress, []*TbAddress) {

	var add []*TbAddress
	var addto []*TbAddress

	for _, v := range active.AddressFrom {
		//1出发地 2目的地
		if v.Type == 1 {
			add = append(add, v)
		} else if v.Type == 2 {
			addto = append(addto, v)
		}
	}

	active.AddressFrom = add
	active.AddressTo = addto

	return active.AddressFrom, active.AddressTo

}

//解耦活动日期格式化
func (TbActivity *TbActivity) FormatTime1(active *TbActivity) (string, string, string, string) {

	active.SignStartTime = active.SignStartTime[:10]
	active.SignEndTime = active.SignEndTime[:10]

	active.ActiveStartTime = active.ActiveStartTime[:10]
	active.ActiveEndTime = active.ActiveEndTime[:10]

	return active.SignStartTime, active.SignEndTime, active.ActiveStartTime, active.ActiveEndTime
}
