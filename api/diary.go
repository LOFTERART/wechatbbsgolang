package api

import (
	"QUZHIYOU/service/home"
	"github.com/gin-gonic/gin"
)

func HomeList(c *gin.Context) {

	//获取header传参
	userId := c.Request.Header.Get("userId")

	var paramsHome home.ListDiaryService

	if err := c.ShouldBind(&paramsHome); err == nil {
		res := paramsHome.GetDiarys(userId)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}


}
