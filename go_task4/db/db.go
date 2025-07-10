package db

import (
	"github.com/zzp326612343/go_study/go_task4/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/task4?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败:" + err.Error())
	}

	return db
}

func CreateTable() {
	db := InitDB()
	db.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})
}
