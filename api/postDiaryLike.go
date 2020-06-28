package api

import (
	"QUZHIYOU/services/home"
	"github.com/gin-gonic/gin"
)

func PostDiaryLike(c *gin.Context) {

	userId := c.Request.Header.Get("userId")

	var userids []string
	userids = append(userids, userId)

	var PassUserId home.DiaryLikeService

	if err := c.ShouldBind(&PassUserId); err == nil {
		res := PassUserId.LikeDiary(userids)
		c.JSON(200, &res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}
