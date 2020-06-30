package api

import (
	"QUZHIYOU/models"
	"QUZHIYOU/serializer"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strings"
)



type FollowUsers struct {
	UserId uint `json:"user_id"`
	Type    bool `json:"type" form:"type"`
}

func (item *FollowUsers) Follow(userid []string) serializer.Response {
	
	//找到我的用户信息
	myUserInfo := models.User{
		Model: gorm.Model{
			ID: item.UserId,
		},
	}
	models.DB.First(&myUserInfo)

	//在我的用户信息里面更新关注人id

	UserLikeArrayId := strings.Split(myUserInfo.Follow, ",")


	if item.Type {
		for k, v := range UserLikeArrayId {
			if userid[0] == v {
				UserLikeArrayId = append(UserLikeArrayId[:k], UserLikeArrayId[k+1:]...)
			}
		}

		str := strings.Join(UserLikeArrayId, ",")

		myUserInfo.Follow = str

		models.DB.Model(&myUserInfo).
			Updates(map[string]interface{}{"follow": myUserInfo.Follow})

	} else {
		for _, v := range UserLikeArrayId {
			if userid[0] != v {
				userid = append(userid, v)
			}
		}
		str := strings.Join(userid, ",")

		myUserInfo.Follow = str

		models.DB.Model(&myUserInfo).
			Updates(map[string]interface{}{ "follow": myUserInfo.Follow})
	}


	return serializer.Response{
		Code:  0,
		Data:  nil,
		Msg:   "操作成功",
	}

}

//关注某个用户 传递关注用户ID
func FollowUser(c *gin.Context) {
	userId := c.Request.Header.Get("userId")
	var info FollowUsers
	if err := c.ShouldBind(&info); err == nil {
		var userids []string  //关注用户id
		userids = append(userids, userId)
		res := info.Follow(userids)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}
