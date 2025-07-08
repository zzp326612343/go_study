package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Age   int
	Grade string
}

var db *sqlx.DB

type Employee struct {
	ID         int     `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

func initDB(dsn string) {
	var err error
	db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Println("数据库连接失败:", err)
	}
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败:" + err.Error())
	}
	// OperateAccount(db)
	// initDB(dsn)
	// var employees []Employee
	// _, err := db.NamedExec("insert into employees (name,department,salary) values (:name,:department,:salary)", Employee{Name: "杂21", Department: "产品部", Salary: 13500})
	// _, err := db.Exec("insert into employees (name,department,salary) values (?,?,?)", "杂34", "产品部", 13500)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// db.Select(&employees, "select * from employees where department = ?", "技术部")
	// for _, e := range employees {
	// 	fmt.Println(e)
	// }
	// var emp Employee
	// db.Get(&emp, "select * from employees order by salary desc limit 1")
	// fmt.Println(emp)

	createUserPostComment(db)
	// queryUserPostsWithComments(db, 1)
	// findMostCommentedPost(db)
}

func operateStudent(db *gorm.DB) {
	db.AutoMigrate(&Student{})

	// db.Create(&Student{Name: "张三2", Age: 14, Grade: "三年级"})
	var stus []Student
	if err := db.Where("age > ?", 18).Find(&stus).Error; err != nil {
		panic("查询失败: " + err.Error())
	}
	// db.First(&stu)
	db.Find(&stus)
	fmt.Println(stus)
	// var st Student
	// db.Model(&st).Where("name = ?", "张三").Update("grade", "四年级")
	// fmt.Println(st)
	// db.Where("age < ?", 15).Delete(&Student{})
}

type Account struct {
	ID      int `gorm:"primaryKey"`
	Balance float64
}

type Transaction struct {
	ID            int `gorm:"primaryKey"`
	FromAccountId int
	Amount        float64
}

func OperateAccount(db *gorm.DB) {
	db.AutoMigrate(&Account{}, &Transaction{})
	err := db.Transaction(func(tx *gorm.DB) error {
		amount := 100.00
		if err := tx.Create(&[]Account{{Balance: 200}, {Balance: 100}}).Error; err != nil {
			return err
		}
		var list []Account
		tx.Find(&list)
		fmt.Println(list)
		var A Account
		tx.Where("id =?", 1).Find(&A)
		if A.Balance < amount {
			fmt.Println("余额不够")
			return nil
		}
		if err := tx.Model(&Account{}).Where("id=?", A.ID).Update("balance", A.Balance-amount).Error; err != nil {
			return err
		}
		if err := tx.Model(&Account{}).Where("id=?", 2).Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
			return err
		}
		fmt.Println("事务成功")
		tx.Find(&list)
		fmt.Println(list)
		return nil
	})
	if err != nil {
		fmt.Println("事务失败:", err)
	}
}

func OperateAccount2(db *gorm.DB) {
	tx := db.Begin()
	tx.Rollback()
	tx.Commit()
}

type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(255)"`
	PostCount int
	Posts     []Post
}

type Post struct {
	gorm.Model
	Title         string `gorm:"type:varchar(255)"`
	UserID        uint
	User          User `gorm:"foreignKey:UserID"`
	Comments      []Comment
	CommentStatus string `gorm:"type:varchar(255)"`
}

type Comment struct {
	gorm.Model
	Content string `gorm:"type:varchar(255)"`
	PostID  uint
	Post    Post `gorm:"foreignKey:PostID"`
}

func createUserPostComment(db *gorm.DB) {
	err := db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		fmt.Println("创建失败")
	}
	// 创建示例数据
	// user := User{Name: "张三"}
	// db.Create(&user)

	// post1 := Post{Title: "第一篇文章", UserID: user.ID}
	// post2 := Post{Title: "第二篇文章", UserID: user.ID}
	// db.Create(&post1)
	// db.Create(&post2)

	// comment1 := Comment{Content: "不错", PostID: post1.ID}
	// comment2 := Comment{Content: "写得很好", PostID: post1.ID}
	// comment3 := Comment{Content: "继续加油", PostID: post2.ID}
	// db.Create(&comment1)
	// db.Create(&comment2)
	// db.Create(&comment3)

	// post2 := Post{Title: "第SAN篇文章", UserID: 1}
	// db.Create(&post2)
	fmt.Println("即将删除评论")
	var c Comment
	if err := db.First(&c, 3).Error; err != nil {
		fmt.Println("评论不存在")
		return
	}
	db.Delete(&c, 3)
	fmt.Println("已请求删除评论")
}

func queryUserPostsWithComments(db *gorm.DB, userId int) {
	var user User
	err := db.Model(&User{}).Preload("Posts.Comments").First(&user, userId).Error
	if err != nil {
		panic("查询失败: " + err.Error())
	}
	fmt.Printf("用户：%s\n", user.Name)
	for _, post := range user.Posts {
		fmt.Printf("  文章：%s\n", post.Title)
		for _, comment := range post.Comments {
			fmt.Printf("    评论：%s\n", comment.Content)
		}
	}
}

type PostWithCount struct {
	PostID       uint
	Title        string
	CommentCount int
}

func findMostCommentedPost(db *gorm.DB) {
	var result PostWithCount
	err := db.Table("posts").
		Select("posts.id as post_id, posts.title, COUNT(comments.id) as comment_count").
		Joins("left join comments on comments.post_id = posts.id").
		Group("posts.id").
		Order("comment_count DESC").
		Limit(1).
		Scan(&result).Error
	if err != nil {
		fmt.Println("查询失败:", err)
		return
	}

	fmt.Printf("评论最多的文章：ID=%d，标题=%s，评论数=%d\n", result.PostID, result.Title, result.CommentCount)

}

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	err = tx.Model(&User{}).Where("id =?", p.UserID).Update("post_count", gorm.Expr("post_count + ?", 1)).Error
	if err != nil {
		return fmt.Errorf("更新用户文章数量失败: %v", err)
	}
	return nil
}

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	fmt.Println("⚠️ 已删除评论 ID:", c.ID)

	var count int64
	if err := tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count).Error; err != nil {
		return err
	}

	fmt.Println("剩余评论数：", count)

	if count == 0 {
		err := tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_status", "无评论").Error
		if err != nil {
			return err
		}
		fmt.Println("✅ 已更新文章为无评论状态")
	}
	return nil
}
