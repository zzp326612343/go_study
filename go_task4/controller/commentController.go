package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/zzp326612343/go_study/go_task4/db"
	"github.com/zzp326612343/go_study/go_task4/model"
)

type Comment struct {
	Content string `json:"content" form:"content"`
	PostID  uint   `json:"postId" form:"postId"`
	UserID  uint   `json:"userId" form:"userId"`
}

func CreateComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var comment Comment
		if err := c.ShouldBind(&comment); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		val, ok := c.MustGet("userId").(uint)
		if !ok {
			c.JSON(400, gin.H{"error": "invalid user id type"})
			return
		}
		db := db.InitDB()
		if err := db.Create(&model.Comment{Content: comment.Content, PostID: comment.PostID, UserID: val}).Error; err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "comment created"})
	}
}

func GetComments() gin.HandlerFunc {
	return func(c *gin.Context) {
		var comment Comment
		c.ShouldBindQuery(&comment)
		fmt.Printf("绑定参数: %+v\n", comment)
		db := db.InitDB()
		var comments []model.Comment
		query := db.Model(&model.Comment{}) // 构造查询链

		if comment.PostID != 0 {
			query = query.Where("post_id = ?", comment.PostID)
		}
		if comment.UserID != 0 {
			query = query.Where("user_id = ?", comment.UserID)
		}
		if comment.Content != "" {
			query = query.Where("content LIKE ?", "%"+comment.Content+"%")
		}

		if err := query.Find(&comments).Error; err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, comments)
	}
}
