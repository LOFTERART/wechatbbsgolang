package api

import (
	"QUZHIYOU/services/wxLogin"
	"github.com/gin-gonic/gin"
	"strconv"
)

func WxLoginUserInfo(c *gin.Context)  {

	var userinfo wxLogin.WxUserInfo

	userid:=c.Request.Header.Get("userId")

	i,_:=strconv.Atoi(userid)

	if err:=c.ShouldBind(&userinfo);err==nil{
		res:=userinfo.GetUserInfo(uint(i))
		c.JSON(200,&res)
	}else {
		c.JSON(200,ErrorResponse(err))
	}

}