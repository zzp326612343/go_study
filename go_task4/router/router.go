package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zzp326612343/go_study/go_task4/controller"
	"github.com/zzp326612343/go_study/go_task4/middleware"
)

func UserRouter(c *gin.Engine) {
	c.Use(middleware.LogMiddleware())
	user := c.Group("/user")
	user.POST("/register", controller.Register())
	user.POST("/login", controller.Login())

	api := c.Group("/api")
	api.Use(middleware.JWTAuthMiddleware())
	post := api.Group("/post")
	post.POST("/", controller.CreatePost())
	post.GET("/", controller.GetPosts())
	post.GET("/:id", controller.GetPost())
	post.PUT("/:id", controller.UpdatePost())
	post.DELETE("/:id", controller.DeletePost())
	comment := api.Group("/comment")
	comment.POST("/", controller.CreateComment())
	comment.GET("/", controller.GetComments())
}
