package serializer

import (
	"QUZHIYOU/middleware"
	"QUZHIYOU/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type User struct {
	Id    uint   `  json:"userId" `
	Token string `json:"token"`
	NickName  string `json:"nick_name"`
	AvatarUrl string `json:"avatar_url"`
	Province  string `json:"province"`
	City      string `json:"city"`
	Country   string `json:"country"`
	Gender    int    `json:"gender"`
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

func BuildUserFormat(item models.User) User {
	return User{
		Id:        item.ID,
		NickName:  item.NickName,
		AvatarUrl: item.AvatarUrl,
		Province:  item.Province,
		City:      item.City,
		Country:   item.Country,
		Gender:    item.Gender,
	}
}