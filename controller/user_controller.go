package controller

import (
	"4096Tiktok/dao"
	"4096Tiktok/service"
	//"fmt"
	"github.com/gin-gonic/gin"

	"net/http"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
var usersLoginInfo = map[string]User{
	"zhangleidouyin": {
		Id:            1,
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

var userIdSequence = int64(1)

type UserLoginResponse struct {
	Response
	UserId int  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User User `json:"user"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	if service.VerifyNameAndPwd(username, password) != true {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response:Response{StatusCode: 201, StatusMsg: "invalid username or password"},
		})
		return
	}

	encryptedPwd := service.EncryptPwd(password)
	user := dao.User{Username: username, Password: encryptedPwd}
	if err := service.AddUser(&user); err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response:Response{StatusCode: 202, StatusMsg: "username registered"},
		})
	}else {
		token, _ := service.ReleaseToken(&user)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId: int(user.ID),
			Token:    token,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	if user, err := service.GetUserByName(username); err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response:Response{StatusCode: 203, StatusMsg: "user doesn't exist"},
		})
	}else {
		if service.DecryptPwd(password, user.Password) != true {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response:Response{StatusCode: 204, StatusMsg: "incorrect password"},
			})
		}else {
			token, _ := service.ReleaseToken(&user)
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 0},
				UserId: int(user.ID),
				Token:    token,
			})
		}
	}

}

func UserInfo(c *gin.Context) {
	//token := c.Query("token")
	//user_id := c.Query("user_id")

	//user := dao.User{}

	//user, _ := c.Get("user")
	//User := user.(dao.User)
	////Id := User.ID
	//fmt.Println("user: ", user)

	//if user, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0, StatusMsg: "test ok"},
			//User:     user,
		})
	//} else {
	//	c.JSON(http.StatusOK, UserResponse{
	//		Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
	//	})
	//}
}
