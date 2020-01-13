package api

import (
	"QUZHIYOU/services/home"
	"github.com/gin-gonic/gin"
)

func GetDiaryInfo(c *gin.Context)  {

	var diary home.DiaryInfoService
	res:=diary.GetDiaryInfo()
	c.IndentedJSON(200,&res)


}