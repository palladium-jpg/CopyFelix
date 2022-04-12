package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var r *gin.Engine

func main() {
	//fmt.Println("Hello World!!!")
	//db := InitDb()
	//defer db.Close()
	//db = common.InitDb()
	r = gin.Default()
	r = CollectRoute(r)
	err := r.Run()
	if err != nil {
		return
	}

}
