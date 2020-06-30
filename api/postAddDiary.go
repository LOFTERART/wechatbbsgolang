package api

import (
	"QUZHIYOU/service/home"
	"github.com/gin-gonic/gin"
	"strconv"
)

func PostAddDiary(c *gin.Context) {

	userid := c.Request.Header.Get("userId")
	i, _ := strconv.Atoi(userid)

	var dia home.AddDiaryService

	if err := c.ShouldBind(&dia); err == nil {
		res := dia.AddDiary(uint(i))
		c.JSON(200, &res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}
