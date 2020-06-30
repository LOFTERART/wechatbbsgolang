package api

import (
	"QUZHIYOU/service/wxLogin"
	"github.com/gin-gonic/gin"
)

func WxLogin(c *gin.Context) {

	code := wxLogin.UserService{}

	if err := c.ShouldBind(&code); err == nil {
		res := code.WxUserLogin()
		c.JSON(200, &res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}
