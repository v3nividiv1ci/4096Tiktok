package dao

import (
	"gorm.io/gorm"
	"log"
)

type Video struct {
	gorm.Model
	AuthorId int
	PlayUrl string
	CoverUrl string
	Title string
}

func AddVideo(video *Video) error {
	DB := GetDB()
	tx := DB.Begin()
	if err := tx.Model(&Video{}).Create(&video).Error; err != nil {
		tx.Rollback()
		log.Println(err.Error())
		return err
	}
	tx.Commit()
	return nil
}
