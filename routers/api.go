package routers

import (
	"QUZHIYOU/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {

	// 初始化默认路由
	router := gin.Default()
	router.StaticFS("/static", http.Dir("./static"))

	router.GET("/wxlogin", api.WxLogin)
	v1 := router.Group("v1")
	{
		//首页活动列表
		v1.GET("/wechat/activity/selectActivityList", api.ActivityList)
		//活动详情信息
		v1.GET("/wechat/activity/selectActivicyInfoById", api.ActivityInfo)
		//获取code码
		v1.GET("/getqrcode", api.Getqrcode)
	}

	//动态首页

	home := router.Group("home")
	{

		home.GET("/homediarys", api.HomeList)
		home.GET("/classify", api.Classify)
		home.GET("/ad", api.GetAd)
		home.POST("/adddiary", api.PostAddDiary)
		home.POST("/upload", api.PostDiaryPic)
		home.GET("/getcommunity", api.Getcommunity)
	}

	return router

}
