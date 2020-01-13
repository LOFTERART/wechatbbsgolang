package serializer

import "QUZHIYOU/models"

type UserPics struct {
	UserPics string `json:"user_pics"`
}

func BuildHomeDiaryInfoUserPic(item *models.User) *UserPics {

	return &UserPics{UserPics: item.AvatarUrl}

}

func BuildHomeDiaryInfoUserPics(item []*models.User) ( []*UserPics) {

	pics:=make([]*UserPics,0)

	for _, v := range item {
		pics = append(pics, BuildHomeDiaryInfoUserPic(v))
	}

	return pics

}
