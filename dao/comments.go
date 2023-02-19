package dao

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserId int
	VideoId int
	Text string
}
