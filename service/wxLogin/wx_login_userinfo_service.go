package wxLogin

import (
	"QUZHIYOU/models"
	"QUZHIYOU/serializer"
	"github.com/jinzhu/gorm"
)

type WxUserInfo struct {
	NickName  string `json:"nick_name"`
	AvatarUrl string `json:"avatar_url"`
	Province  string `json:"province"`
	City      string `json:"city"`
	Country   string `json:"country"`
	Gender    int    `json:"gender"`
}

func (service *WxUserInfo) GetUserInfo(userid uint) serializer.Response {

	//查找id 对应的userinfo
	user := models.User{
		Model: gorm.Model{
			ID: userid,
		},
	}

	models.DB.First(&user).
		Updates(map[string]interface{}{"nick_name": service.NickName, "avatar_url": service.AvatarUrl, "province": service.Province, "gender": service.Gender, "city": service.City, "country": service.Country})

	return serializer.Response{
		Code:  0,
		Data:  nil,
		Msg:   "更新成功",
		Error: "",
	}

}
