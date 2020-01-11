package api

import (
	"QUZHIYOU/services/home"
	"github.com/gin-gonic/gin"
)

func HomeList(c *gin.Context) {

	var paramsHome home.ListDiaryService

	if err := c.ShouldBind(&paramsHome); err == nil {
		res := paramsHome.GetDiarys()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}


}
