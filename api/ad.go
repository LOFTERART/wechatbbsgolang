package api

import (
	"QUZHIYOU/service/home"
	"github.com/gin-gonic/gin"
)

func GetAd(c *gin.Context)  {

	var ads home.AdService

	res:=ads.GetAds()
	c.JSON(200,&res)

}