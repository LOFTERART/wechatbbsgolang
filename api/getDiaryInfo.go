package api

import (
	"QUZHIYOU/services/home"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetDiaryInfo(c *gin.Context)  {
	userId := c.Request.Header.Get("userId")
	i, _ := strconv.Atoi(userId)
	var Diary home.DiaryInfoService
	if err:=c.ShouldBind(&Diary);err==nil{

		fmt.Println(Diary.Id,"---------diatyid----")
		Res:=Diary.GetDiaryInfo(uint(i))
		c.JSON(200,&Res)
	}



}