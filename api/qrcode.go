package api

import (
	"github.com/gin-gonic/gin"
	"github.com/medivhzhan/weapp/v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	token string
	dir   string
)

func Getqrcode(c *gin.Context) {

	//获取文件路径
	dir, _ = filepath.Abs(filepath.Dir(os.Args[0]))

	token, _ := weapp.GetAccessToken(os.Getenv("WXAPP_ID"), os.Getenv("WXSECRET"))

	getter := weapp.UnlimitedQRCode{
		Scene:     "id=1",
		Page:      "homeSub/pages/homeDetail/homeDetail",
		Width:     430,
		AutoColor: true,
		LineColor: weapp.Color{"255", "255", "255"},
		IsHyaline: false,
	}

	resp, res, err := getter.Get(token.AccessToken)

	if err != nil {
		return
	}
	if err := res.GetResponseError(); err != nil {
		return
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)

	f, _ := os.Create(dir + "/static/" + "1.png")

	defer f.Close()

	f.Write(content)

	//上传图片到服务器 -no

	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "生成成功",
		"data": dir + "/static/" + "1.png",
	})

}
