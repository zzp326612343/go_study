package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zzp326612343/go_study/go_task4/db"
	"github.com/zzp326612343/go_study/go_task4/model"
)

type Post struct {
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	UserID  uint   `json:"userId" form:"userId"`
}

func CreatePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var post Post
		if err := c.ShouldBind(&post); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		db := db.InitDB()
		userId, ok := c.MustGet("userId").(uint)
		if !ok {
			c.JSON(400, gin.H{"error": "invalid user ID type"})
			return
		}
		if err := db.Create(&model.Post{Title: post.Title, Content: post.Content, UserID: userId}).Error; err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		}
		c.JSON(201, gin.H{"message": "post created"})
	}
}

func GetPosts() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := db.InitDB()
		var posts []model.Post
		if err := db.Find(&posts).Error; err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, posts)
	}
}

func GetPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := db.InitDB()
		var post model.Post
		post.Comments = make([]model.Comment, 0)
		if err := db.Preload("Comments").First(&post, c.Param("id")).Error; err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, post)
	}
}

func UpdatePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := db.InitDB()
		var dbpost model.Post
		if err := db.First(&dbpost, c.Param("id")).Error; err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if c.MustGet("userId") != dbpost.UserID {
			c.JSON(400, gin.H{"error": "you are not the author of this post"})
			return
		}
		var post Post
		if err := c.ShouldBind(&post); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if err := db.Model(&model.Post{}).Where("id = ?", c.Param("id")).Update("title", post.Title).Update("content", post.Content).Error; err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "post updated"})
	}
}

func DeletePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := db.InitDB()
		var dbpost model.Post
		if err := db.First(&dbpost, c.Param("id")).Error; err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		if c.MustGet("userId") != dbpost.UserID {
			c.JSON(400, gin.H{"error": "you are not the author of this post"})
			return
		}
		if err := db.Delete(&model.Post{}, c.Param("id")).Error; err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
		}
		c.JSON(200, gin.H{"message": "post deleted"})
	}
}
