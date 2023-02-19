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
	// Todo: find user in the database
	user := User{}
	return user, nil
}






