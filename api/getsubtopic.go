package api

import (
	"QUZHIYOU/service/home"
	"github.com/gin-gonic/gin"
)

func Getsubtopic(c *gin.Context) {

	var subtopic home.SubTopicService
	if err := c.ShouldBind(&subtopic); err == nil {
		res := subtopic.GetSubTopic()
		c.JSON(200, &res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}
