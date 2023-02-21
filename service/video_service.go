package service

import (
	"4096Tiktok/dao"
	"log"
)

func AddVideo(video *dao.Video) error {
	if err := dao.AddVideo(video); err != nil {
		log.Println("AddVideo failure")
		return err
	}
	return nil
}

func GetVideoByUserIDAndTitle(UID uint, title string) bool {
	if video, err := dao.GetVideoByUserAndTitle(UID, title); err == nil && video.UserID != 0{
		return true
	}
	return false
}