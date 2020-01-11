package models

import "github.com/jinzhu/gorm"

type Communitys struct {
	gorm.Model
	Name    string
	KeyWord string
	Letter  string
	IsOpen  bool
}



