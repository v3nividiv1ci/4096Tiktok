package model

import "gorm.io/gorm"

type Like struct {
	gorm.Model
	UserId int
	VideoId int
	cancel bool
}