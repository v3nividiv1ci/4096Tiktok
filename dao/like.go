package dao

import (
	"fmt"
	"gorm.io/gorm"
	"log"
)

func GetVideoById(VideoId int) (Video, error) {
	DB := GetDB()
	tx := DB.Begin()
	var video Video
	if err := tx.Model(&Video{}).Where("video_id = ?", VideoId).First(&video).Error; err != nil {
		tx.Rollback()
		log.Println(err.Error())
		return video, err
	}

	tx.Commit()
	return video, nil
}

func LikeVideo(video *Video, UserId int) error {
	DB := GetDB()
	tx := DB.Begin()
	if err := tx.Model(&video).Association("Likes").Append(&User{UserID: uint(UserId)}); err != nil {
		tx.Rollback()
		log.Println(err.Error())
		return err
	}
	tx.Commit()
	return nil
}

func DislikeVideo(video *Video, UserId int) error {
	DB := GetDB()
	tx := DB.Begin()
	if err := tx.Model(&video).Association("Likes").Delete(&User{UserID: uint(UserId)}); err != nil {
		tx.Rollback()
		log.Println()
		return err
	}
	tx.Commit()
	return nil
}

func VideoLikeCount(VideoId int) int64 {
	var count int64
	DB := GetDB()
	tx := DB.Begin()
	count = tx.Model(&Video{}).Where("video_id = ?", VideoId).Association("Likes").Count()
	tx.Commit()
	return count
}

func UserLikeCount(UserId int) int64 {
	DB := GetDB()
	tx := DB.Begin()
	var user User
	tx.Preload("Likes").First(&user, UserId)
	var likesCount int64
	for _, video := range user.Likes {
		likesCount += tx.Model(&video).Association("Likes").Count()
	}
	tx.Commit()
	return likesCount
}

func GetUserLikedCount(Id int) int64 {
	DB := GetDB()
	tx := DB.Begin()

	var user User
	tx.Preload("Videos.Likes").Find(&user)

	var totalLikes int64
	for _, video := range user.Videos {
		totalLikes += int64(len(video.Likes))
	}
	tx.Commit()
	return totalLikes
}

func UserAllLikedCount(ids [] uint) int64 {
	var count int64
	DB := GetDB()
	tx := DB.Begin()

	fmt.Println("video_ids are: ", ids)
	//count = tx.Model(&Video{}).Where("video_id IN ?", ids).Association("Likes").Count()
	tx.Preload("Videos", func(db *gorm.DB) *gorm.DB {
		return db.Where("video_id in ?", ids)
	}).Count(&count)
	tx.Commit()
	return count
}