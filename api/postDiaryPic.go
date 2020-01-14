package api

import (
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"image"
	"image/color"
	"log"
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
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
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

		src, err := imaging.Open(dir + "/static/" + file.Filename)
		if err != nil {
			log.Fatalf("failed to open image: %v", err)
		}
		src = imaging.Resize(src, 300, 300, imaging.Lanczos)
		dst := imaging.New(300, 300, color.NRGBA{255, 255, 255, 0})
		dst = imaging.Paste(dst, src, image.Pt(0, 0))
		imaging.Save(dst, dir+"/static/"+"b.jpg")

		c.IndentedJSON(200, gin.H{
			"code": 0,
			"data": back,
			"msg":  "ok",
		})
	}

}
