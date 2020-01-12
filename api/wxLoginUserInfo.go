package api

import (
	"QUZHIYOU/services/wxLogin"
	"github.com/gin-gonic/gin"
	"strconv"
)

func WxLoginUserInfo(c *gin.Context) {
	userid := c.Request.Header.Get("userId")
	i, _ := strconv.Atoi(userid)

	var Userinfo wxLogin.WxUserInfo

	if err := c.ShouldBind(&Userinfo); err == nil {
		res := Userinfo.GetUserInfo(uint(i))
		c.JSON(200, &res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}
