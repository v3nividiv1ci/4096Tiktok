package model

import "gorm.io/gorm"

type Follow struct {
	gorm.Model
	UserId int
	FollowerId int
}