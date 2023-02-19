package service

import (
	"4096Tiktok/dao"
	"4096Tiktok/middleware"
	"golang.org/x/crypto/bcrypt"
	"log"
	"regexp"
)

func CheckString(string string) bool {
	if ok, _ := regexp.MatchString("^[\\w_-]{6,32}$", string); !ok {
		return false
	}
	return true
}

func VerifyNameAndPwd(username, password string) bool {
	if err := CheckString(username) && CheckString(password); err != true {
		return false
	}
	return true
}

func ReleaseToken (user *dao.User) (string, error){
	token, err := middleware.TokenRelease(*user)
	return token, err
}

func EncryptPwd (password string) string {
	EncryptedPwd, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(EncryptedPwd)
}

func DecryptPwd (password, encryptedPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encryptedPwd), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func AddUser (user *dao.User) error {
	if err := dao.AddUser(user); err != nil {
		log.Println("AddUser failure")
		return err
	}
	return nil
}

func GetUserByName (username string) (dao.User, error){
	if user, err := dao.GetUserByName(username); err != nil {
		log.Println("Get user failure")
		return dao.User{}, err
	}else {
		return user, nil
	}
}
