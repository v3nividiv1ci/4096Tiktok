package dao

import (
	"github.com/RaymondCode/simple-demo/dao/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var DB *gorm.DB

func InitDB() *gorm.DB {
	dsn := "root:password@tcp/douyin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database.")
	}
	db.AutoMigrate(&model.User{}, &model.Comment{}, &model.Follow{}, &model.Video{}, &model.Like{})
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}