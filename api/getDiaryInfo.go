package api

import (
	"QUZHIYOU/services/home"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetDiaryInfo(c *gin.Context)  {
	userId := c.Request.Header.Get("userId")
	i, _ := strconv.Atoi(userId)
	var Diary home.DiaryInfoService
	if err:=c.ShouldBind(&Diary);err==nil{

		Res:=Diary.GetDiaryInfo(uint(i))
		c.JSON(200,&Res)
	}



}