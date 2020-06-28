package wxLogin

import (
	"QUZHIYOU/models"
	"QUZHIYOU/serializer"
	"github.com/medivhzhan/weapp/v2"
	"os"
)

type UserService struct {
	Code string `form:"code" json:"code"`
}

func (u *UserService) WxUserLogin() serializer.Response {

	var user models.User

	//	获取code 查询数据库 没有需创建用户

	res, err := weapp.Login(os.Getenv("WXAPP_ID"), os.Getenv("WXSECRET"), u.Code)

	if err != nil {
		// 处理一般错误信息
		return serializer.Response{
			Code:  0,
			Data:  nil,
			Msg:   "2",
			Error: "",
		}
	}

	if err := res.GetResponseError(); err != nil {
		// 处理微信返回错误信息
		return serializer.Response{
			Code:  0,
			Data:  nil,
			Msg:   "1",
			Error: "",
		}
	}


	user= models.User{OpenId: res.OpenID}

	if err:=models.DB.Where("open_id=?",res.OpenID).First(&user).Error;err==nil{
		return serializer.Response{
			Code:0,
			Msg:   "已存在",
			Error: "",
			Data:  serializer.BuildUser(&user),
		}
	}else {
		models.DB.Create(&user)

		return serializer.Response{
			Code:0,
			Msg:   "创建成功",
			Error: "",
			Data:  serializer.BuildUser(&user),
		}
	}

}
