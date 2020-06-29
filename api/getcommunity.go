package api

import (
	"QUZHIYOU/service/home"
	"github.com/gin-gonic/gin"
)



func Getcommunity(c *gin.Context)  {


	var comunity home.HomeCommunityService
	res:=comunity.GetCommunity()
	c.JSON(200,&res)

}