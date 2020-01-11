package middleware

import (
	"QUZHIYOU/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/medivhzhan/weapp/v2"
	"os"
)

func WXToken() gin.HandlerFunc {
	return func(c *gin.Context) {

		res, _ := weapp.GetAccessToken(os.Getenv("WXAPP_ID"), os.Getenv("WXSECRET"))

		if err := res.GetResponseError(); err != nil {
			// 处理微信返回错误信息
			return
		}

		fmt.Printf("返回结果: %#v", res)

		token := models.TbWxtoken{AccessToken: res.AccessToken}

		models.DB.Save(&token)

	}
}
