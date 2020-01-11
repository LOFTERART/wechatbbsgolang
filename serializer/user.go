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
