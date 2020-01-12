package api

import (
	"QUZHIYOU/services/home"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"strconv"
)

func PostDiaryLike(c *gin.Context) {

	userId := c.Request.Header.Get("userId")
	i, _ := strconv.Atoi(userId)

	var userids pq.Int64Array
	userids = append(userids, int64(i))

	var PassUserId home.DiaryLikeService

	if err := c.ShouldBind(&PassUserId); err == nil {
		res := PassUserId.LikeDiary(userids)
		c.JSON(200, &res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}
