package api

import (
	"QUZHIYOU/utils"
	"github.com/gin-gonic/gin"
	"path"
	"strings"
)

//返回字段
type Back struct {
	Url           string `json:"url"`
	PostIndexSort int    `json:"post_indexSort"`
}

func PostDiaryPic(c *gin.Context) {
	//获取文件路径
	//dir, _ := os.Getwd()
	form, _ := c.MultipartForm()
	files := form.File["file"]
	var back Back
	var name string
	for index, file := range files {
		ext := strings.ToLower(path.Ext(file.Filename))
		name = utils.RandStringRunes(10) + ext
		back.Url = name
		back.PostIndexSort = index

		// 上传文件至指定目录
		if err := c.SaveUploadedFile(file, "./static/"+name); err != nil {
			c.JSON(200, gin.H{
				"code": 404,
				"data": nil,
				"msg":  "上传失败",
			})
			return
		}

		c.JSON(200, gin.H{
			"code": 0,
			"data": back,
			"msg":  "ok",
		})
	}


}
