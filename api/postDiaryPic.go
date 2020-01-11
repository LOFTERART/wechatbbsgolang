package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path/filepath"
)

type Back struct {
	Url           string `json:"url"`
	PostIndexSort int `json:"post_indexSort"`
}

func PostDiaryPic(c *gin.Context) {
	//获取文件路径
	dir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
	form, _ := c.MultipartForm()
	files := form.File["file"]

	var back Back
	for index, file := range files {

		back.Url = file.Filename
		back.PostIndexSort=index
		log.Println(file.Filename)
		// 上传文件至指定目录
		c.SaveUploadedFile(file, dir+"/static/"+file.Filename)

		c.IndentedJSON(200, gin.H{
			"code": 0,
			"data": back,
			"msg":  "ok",
		})
	}

}
