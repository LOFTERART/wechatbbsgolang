package serializer

import (
	"QUZHIYOU/middleware"
	"QUZHIYOU/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type User struct {
	Id        uint   `  json:"userId" `
	Token     string `json:"token,omitempty"`
	NickName  string `json:"nick_name,omitempty"`
	AvatarUrl string `json:"avatar_url,omitempty"`
	Province  string `json:"province,omitempty"`
	City      string `json:"city,omitempty"`
	Country   string `json:"country,omitempty"`
	Gender    int    `json:"gender,omitempty"`

	IsFollow              bool      `json:"isFollow"`
	FollowDes string `json:"followDes"`

	Auth                 string `json:"auth,omitempty"`                   //职业认证
	Specialist           bool   `json:"specialist,omitempty"`             //是否认证
	Authentication       bool   `json:"authentication,omitempty"`         //认证
	AuthenticationName   string `json:"authentication_name,omitempty"`    //认证的称号 教师 医生
	IsShowAuthentication bool   `json:"is_show_authentication,omitempty"` //是否显示认证
	Job                  string `json:"job,omitempty"`
	CommunityName        string `json:"community_name,omitempty"`
}

func BuildUser(user *models.User) *User {

	j := &middleware.JWT{
		[]byte("admin"),
	}

	claims := middleware.CustomClaims{
		ID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(10 * time.Second).Unix(), //time.Now().Add(24 * time.Hour).Unix() //15000
			Issuer:    "admin",
		},
	}

	token, _ := j.CreateToken(claims)

	_, err := j.ParseToken(token)

	if err != nil {
		if err == middleware.TokenExpired {
			newToken, err := j.RefreshToken(token)
			if err != nil {
				return &User{
					Id:    user.ID,
					Token: newToken,
				}
			} else {
				return &User{
					Id:    user.ID,
					Token: newToken,
				}

			}
		} else {
			return &User{
				Id:    user.ID,
				Token: token,
			}
		}
	} else {
		return &User{
			Id:    user.ID,
			Token: token,
		}
	}

}

func BuildUserFormat(item *models.User,id string) *User {
	var name string
	if item.UserIsLike(id){
		name="已关注"
	}else {
		name="关注"
	}

	return &User{
		Id:                   item.ID,
		IsFollow:               item.UserIsLike(id),
		FollowDes:               name,
		NickName:             item.NickName,
		AvatarUrl:            item.AvatarUrl,
		Province:             item.Province,
		City:                 item.City,
		Country:              item.Country,
		Gender:               item.Gender,
		Auth:                 item.Auth,
		Specialist:           item.Specialist,
		Authentication:       item.Authentication,
		AuthenticationName:   item.AuthenticationName,
		IsShowAuthentication: item.IsShowAuthentication,
		Job:                  item.Job,
		CommunityName:        item.CommunityName,
	}
}

func BuildUserSFormat(item []*models.User,id string) (items []*User) {

	for _, v := range item {
		items = append(items, BuildUserFormat(v,id))
	}
	return
}
