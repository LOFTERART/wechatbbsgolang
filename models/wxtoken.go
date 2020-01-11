package models

//Token
type TbWxtoken struct {
	Id          string `  gorm:"primary_key;column:id" json:"id"`
	AccessToken string `  gorm:"column:ACCESS_TOKEN" json:"accessToken"`
}
