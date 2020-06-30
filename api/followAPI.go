package api

import (
	"QUZHIYOU/models"
	"QUZHIYOU/serializer"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
	"strings"
)



type FollowUsers struct {
	MyUserId uint `json:"my_user_id"`
	YouUserId uint `json:"you_user_id"`
	Type    bool `json:"type" form:"type"`
}

func (item *FollowUsers) Follow(my []string,you []string) serializer.Response {
	

	MyUserInfo := models.User{
		Model: gorm.Model{
			ID: item.MyUserId,
		},
	}
	models.DB.First(&MyUserInfo)

	//在我的用户信息里面更新关注人id
	MyFollow := strings.Split(MyUserInfo.MyFollow, ",")



	YouUserInfo := models.User{
		Model: gorm.Model{
			ID: item.YouUserId,
		},
	}
	models.DB.First(&YouUserInfo)

	//在我的用户信息里面更新关注人id
	YouFollow := strings.Split(YouUserInfo.Follow, ",")
	


	if item.Type {

		//关注人的列表信息更新 都谁关注他了 把关注人的ID存到他的列表
		for k, v := range MyFollow {
			if my[0] == v {
				MyFollow = append(MyFollow[:k], MyFollow[k+1:]...)
			}
		}

		str := strings.Join(MyFollow, ",")

		MyUserInfo.Follow = str

		models.DB.Model(&MyUserInfo).Updates(map[string]interface{}{"my_follow": MyUserInfo.Follow})

	} else {

		for _, v := range MyFollow {
			if you[0] != v {
				you = append(you, v)
			}
		}
		youUser := strings.Join(you, ",")
		models.DB.Model(&MyUserInfo).Updates(map[string]interface{}{ "my_follow": youUser})



		for _, v := range YouFollow {
			if my[0] != v {
				my = append(my, v)
			}
		}
		myUser := strings.Join(my, ",")
		models.DB.Model(&YouUserInfo).Updates(map[string]interface{}{ "follow": myUser})


	}


	return serializer.Response{
		Code:  0,
		Data:  nil,
		Msg:   "操作成功",
	}

}

//关注某个用户 传递关注用户ID
func FollowUser(c *gin.Context) {
	var info FollowUsers
	if err := c.ShouldBind(&info); err == nil {
		var my []string
		my = append(my, strconv.FormatInt(int64(info.MyUserId),10))
		var you []string
		you= append(you, strconv.FormatInt(int64(info.YouUserId),10))
		
		res := info.Follow(my,you)
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}
