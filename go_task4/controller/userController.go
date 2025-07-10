package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zzp326612343/go_study/go_task4/db"
	"github.com/zzp326612343/go_study/go_task4/model"
	"github.com/zzp326612343/go_study/go_task4/utils"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserName string `form:"userName" json:"userName"`
	Password string `form:"password" json:"password"`
	Email    string `form:"email" json:"email"`
}

func Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusOK, gin.H{"message": "参数错误"})
			return
		}
		pas := ""
		var err error
		if pas, err = bcryptPassword(user.Password); err != nil {
			c.JSON(http.StatusOK, gin.H{"message": "密码错误"})
			return
		}
		db := db.InitDB()
		dbUser := model.User{
			UserName: user.UserName,
			Password: pas,
			Email:    user.Email,
		}
		if err = db.Create(&dbUser).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{"message": "注册失败", "error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user User
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusOK, gin.H{"message": "参数错误"})
			return
		}
		db := db.InitDB()
		dbUser := model.User{}
		db.Where("user_name = ?", user.UserName).First(&dbUser)
		if dbUser.ID == 0 {
			c.JSON(http.StatusOK, gin.H{"message": "用户不存在"})
			return
		}
		if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
			c.JSON(http.StatusOK, gin.H{"message": "密码错误"})
			return
		}
		token, err := utils.GenerateToken(dbUser.UserName, dbUser.ID)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"message": "生成token失败"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "登陆成功", "token": token})
	}
}

/**
* 密码加密
 */
func bcryptPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
