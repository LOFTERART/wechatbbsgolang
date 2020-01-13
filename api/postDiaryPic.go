package api

import (
	"github.com/gin-gonic/gin"
	"os"
	"path/filepath"
)

//返回字段
type Back struct {
	Url           string `json:"url"`
	PostIndexSort int    `json:"post_indexSort"`
}

func PostDiaryPic(c *gin.Context) {
	//获取文件路径
	dir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	form, _ := c.MultipartForm()
	files := form.File["file"]
	var back Back
	for index, file := range files {

		back.Url = file.Filename
		back.PostIndexSort = index
		// 上传文件至指定目录
		if err := c.SaveUploadedFile(file, dir+"/static/"+file.Filename); err != nil {
			c.IndentedJSON(200, gin.H{
				"code": 404,
				"data": nil,
				"msg":  "上传失败",
			})
			return
		}

		c.IndentedJSON(200, gin.H{
			"code": 0,
			"data": back,
			"msg":  "ok",
		})
	}

}
