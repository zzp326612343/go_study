package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zzp326612343/go_study/go_task4/db"
	"github.com/zzp326612343/go_study/go_task4/router"
)

func main() {
	db.CreateTable()
	r := gin.Default()
	router.UserRouter(r)
	r.Run()
}
