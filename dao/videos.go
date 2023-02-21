package dao

import (
	"log"
)

type Video struct {
	VideoID		uint 		`gorm:"primarykey"`
	UserID 		uint
	PlayUrl 	string		`gorm:"not null"`
	CoverUrl	string		`gorm:"not null"`
	Title 		string		`gorm:"not null"`
	Comments	[] *Comment
	Likes		[] *User	`gorm:"many2many:like;joinForeignKey:video_id;joinReferences:user_id;"`
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

func GetVideoByUserAndTitle(UID uint, title string) (Video, error) {
	DB := GetDB()
	tx := DB.Begin()
	video := Video{}
	if err := tx.Where("user_id = ? AND title = ?", UID, title).First(&video).Error; err != nil {
		tx.Rollback()
		log.Println(err.Error())
		return video, err
	}
	tx.Commit()
	return video, nil
}
