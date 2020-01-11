package api

import (
	"QUZHIYOU/services/home"
	"github.com/gin-gonic/gin"
)

func Getsubtopic(c *gin.Context)  {

	var subtopic home.SubTopicService
	if err:=c.ShouldBind(&subtopic);err==nil{
		res:=subtopic.GetSubTopic()
		c.IndentedJSON(200,&res)
	}else {
		c.JSON(200, ErrorResponse(err))
	}





}