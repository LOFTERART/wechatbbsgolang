package api

import (
	"QUZHIYOU/services/activity"
	"github.com/gin-gonic/gin"
)

//获取首页列表

func ActivityList2(this *gin.Context) {

	resch := make(chan interface{}, 10)

	service := activity.ListActivityService{}

	if err := this.ShouldBind(&service); err == nil {

		go func() {
			res := service.List()
			resch <- res
			close(resch)
		}()

		this.JSON(200, <-resch)

	} else {
		this.JSON(200, ErrorResponse(err))
	}

}

//获取活动详情信息
func ActivityInfo(this *gin.Context) {

	service := activity.ActivityInfo{}

	if err := this.ShouldBind(&service); err == nil {
		res := service.GetActivityInfo()
		this.JSON(200, res)
	} else {
		this.JSON(200, ErrorResponse(err))
	}

}

func ActivityList(this *gin.Context) {
	service := activity.ListActivityService{}

	if err := this.ShouldBind(&service); err == nil {
		res := service.List()
		this.IndentedJSON(200, res)
	} else {
		this.JSON(200, ErrorResponse(err))
	}

}
