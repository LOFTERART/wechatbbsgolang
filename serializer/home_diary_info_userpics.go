package serializer

import "QUZHIYOU/models"

type UserPics struct {
	UserPics string `json:"user_pics"`
}

func BuildHomeDiaryInfoUserPic(item *models.User) *UserPics {

	return &UserPics{UserPics: item.AvatarUrl}

}


func BuildHomeDiaryInfoUserPics(item []*models.User) (pics []*UserPics) {

	for _,v:=range item{
		pics = append(pics, BuildHomeDiaryInfoUserPic(v))
	}

	return

}