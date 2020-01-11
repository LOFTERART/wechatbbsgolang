package models

import "github.com/jinzhu/gorm"

type Ad struct {
	gorm.Model
	Name string
	Type string
	Link string
	IsShelves bool
	Image string
}
