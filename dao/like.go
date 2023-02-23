package dao

import "log"

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