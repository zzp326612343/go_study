package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `gorm:"type:varchar(255);not null;unique"`
	Password string `gorm:"type:varchar(255)"`
	Email    string `gorm:"type:varchar(255);not null;unique"`
	Posts    []Post
	Comments []Comment
}

type Post struct {
	gorm.Model
	Title    string `gorm:"type:varchar(255)"`
	Content  string `gorm:"type:text"`
	UserID   uint   `gorm:"foreignKey"`
	Comments []Comment
}

type Comment struct {
	gorm.Model
	Content string `gorm:"type:text"`
	UserID  uint   `gorm:"foreignKey"`
	PostID  uint   `gorm:"foreignKey"`
}
