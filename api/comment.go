package api

import (
	service "QUZHIYOU/services"
	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context)  {

	info := service.Comment{}

	if err := c.ShouldBind(&info); err != nil {
		c.JSON(201, ErrorResponse(err))
	} else {
		res := info.CreateComment()
		c.JSON(200, &res)
	}

}
