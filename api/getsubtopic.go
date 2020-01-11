package api

import (
	"QUZHIYOU/services/home"
	"github.com/gin-gonic/gin"
)

func Getsubtopic(c *gin.Context)  {

	var subtopic home.SubTopicService
	res:=subtopic.GetSubTopic()
	c.IndentedJSON(200,&res)

}