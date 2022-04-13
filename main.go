package main

import (
	"CopyFelix/common"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
	"os"
	"xorm.io/xorm"
)

//当在代码中使用了第三方库 ，但是go.mod中并没有跟着更新的时候
//
//如果直接run或者build就会报 missing go.sum entry...
//使用 go mod tidy命令来管理依赖 删除不需要的依赖包，下载新的依赖包，更新go.mod
var r *gin.Engine

func main() {
	//fmt.Println("Hello World!!!")
	//db := InitDb()
	//defer db.Close()
	InitConfig()
	//fmt.Println(viper.Get("datasource.driverName"))
	db := common.InitDb()
	defer func(db *xorm.Engine) {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}(db)
	//defer 函数会将其后面跟随的语句进行延迟处理，先被defer的语句将在最后处理，同理可带入多个defer申请
	r = gin.Default()
	r = CollectRoute(r)
	err := r.Run()
	if err != nil {
		return
	}

}

func InitConfig() {
	workDir, _ := os.Getwd()
	//fmt.Printf(workDir)
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "\\config")
	err := viper.ReadInConfig()
	if err != nil {
		panic("")
	}
}
