package api

import (
	"QUZHIYOU/services/home"
	"github.com/gin-gonic/gin"
)

func PostAddDiary(c *gin.Context)  {

	var dia home.AddDiaryService

	if err:=c.ShouldBind(&dia);err==nil{
		res:=dia.AddDiary()
		c.JSON(200,&res)
	}else {
		c.JSON(200, ErrorResponse(err))
	}

}
