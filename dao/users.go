package dao

import (
	"log"
)

type User struct {
	UserID 		uint			`gorm:"primarykey"`
	Username 	string 		`gorm:"index:,unique"`
	Password 	string
	Videos		[] *Video
	Comments	[] *Comment
	Likes		[] *Video 	`gorm:"many2many:like;joinForeignKey:user_id;joinReferences:video_id;"`
	Fans		[] *User 	`gorm:"many2many:follow;joinForeignKey:user_id;joinReferences:fan_id;"`
}

func AddUser(user *User) error{
	DB := GetDB()
	tx := DB.Begin()
	if err := tx.Model(&User{}).Create(&user).Error; err != nil {
		tx.Rollback()
		log.Println(err.Error())
		return err
	}
	tx.Commit()
	return nil
}

func GetUserByName(name string) (User, error) {
	DB := GetDB()
	tx := DB.Begin()
	user := User{}
	if err := tx.Where("username = ?", name).First(&user).Error; err != nil {
		tx.Rollback()
		log.Println(err.Error())
		return user, err
	}
	tx.Commit()
	return user, nil
}








