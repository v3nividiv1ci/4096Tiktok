package dao

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	AuthorId int
	PlayUrl string
	CoverUrl string
	Title string
}
