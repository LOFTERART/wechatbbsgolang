package api

import (
	service "QUZHIYOU/service"
	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {

	info := service.Comment{}

	if err := c.ShouldBind(&info); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {
		res := info.CreateComment()
		c.JSON(200, &res)
	}

}

func GetComment(c *gin.Context) {

	userID := c.Request.Header.Get("userId") //获取请求的USER—iD

	info := service.Comment{}

	if err := c.ShouldBind(&info); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {
		res := info.GetALLComment(userID)
		c.JSON(200, &res)
	}
}

//点赞vs取消点赞
func LikeDiaryComment(c *gin.Context) {

	userId := c.Request.Header.Get("userId")

	var userids []string
	userids = append(userids, userId)

	var PassUserId service.LikeDairyComment

	if err := c.ShouldBind(&PassUserId); err == nil {
		res := PassUserId.LikeComment(userids)
		c.JSON(200, &res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}
