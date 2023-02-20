package controller

import (
	"4096Tiktok/dao"
	"4096Tiktok/ossDB"
	"4096Tiktok/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	title := c.Query("title")
	fileHeader, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 301,
			StatusMsg:  err.Error(),
		})
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 302,
			StatusMsg:  err.Error(),
		})
		return
	}

	User, _ := c.Get("user")
	user := User.(dao.User)
	filename := filepath.Base(fileHeader.Filename)

	playName := fmt.Sprintf("video/%d_%s", user.ID, filename)
	if err := oss.PutObject(playName, file); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 302,
			StatusMsg:  err.Error(),
		})
		return
	}

	//CoverName := fmt.Sprintf("cover/%d_%s", user.ID, filename)
	playUrl := oss.GeneratePlayUrl(playName)
	coverUrl := oss.GenerateCoverUrl(playUrl)

	video := dao.Video{
		AuthorId: int(user.ID),
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
		Title:    title,
	}
	if err = service.AddVideo(&video); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 303,
			StatusMsg:  err.Error(),
		})
	}

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  playName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}