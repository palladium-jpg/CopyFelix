package main

import (
	"CopyFelix/common"
	"CopyFelix/model"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	//fmt.Println("Hello World!!!")
	//db := InitDb()
	//defer db.Close()
	db := common.GetDB()
	er := db.Sync2(new(model.User))
	if er != nil {
		log.Default()
	}
	r := gin.Default()
	r = CollectRoute(r)
	panic(r.Run())

}
