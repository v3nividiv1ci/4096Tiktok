package service

import "4096Tiktok/dao"

func GetVideoById(VideoId int) (dao.Video, error) {
	var video dao.Video
	if video, err := dao.GetVideoById(VideoId); err == nil && video.UserID != 0 {
		return video, err
	}
	return video, nil
}

func FavorVideo(UserId int, action int, video *dao.Video) error {
	var err error
	switch action {
		case 1:
			err = dao.LikeVideo(video, UserId)
		case 2:
			err = dao.DislikeVideo(video, UserId)
	}
	return err
}