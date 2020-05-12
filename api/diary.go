package api

import (
	"QUZHIYOU/services/home"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func HomeList(c *gin.Context) {

	userid := c.Request.Header.Get("userId")
	fmt.Println(userid,"uid---")
	i, _ := strconv.Atoi(userid)

	var paramsHome home.ListDiaryService

	if err := c.ShouldBind(&paramsHome); err == nil {
		res := paramsHome.GetDiarys(int64(i))
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}


}
