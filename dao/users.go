package dao

import (
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"index:,unique"`
	Password string
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








